package cache

import (
	"SimpleServer/internal/models"
	"SimpleServer/internal/providers/db"
	"fmt"
	"log"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	items             map[string]*models.Item
	defaultExpiration time.Duration
	cleanUpInterval   time.Duration
	db                *db.DataBase
}

// constructor for cache

func NewCache(defaultExpiration, cleanupInterval time.Duration, db *db.DataBase) *Cache {

	// initializing map
	items := make(map[string]*models.Item)
	cache := Cache{

		items:             items,
		defaultExpiration: defaultExpiration,
		cleanUpInterval:   cleanupInterval,
		db:                db,
	}

	// starting gc
	go cache.garbageCollector()

	return &cache
}

// deleting items from cache

func (c *Cache) clearItems(keys []string) {
	c.Lock()

	defer c.Unlock()
	for _, key := range keys {
		delete(c.items, key)
	}

}

// deleting employee from cache and db

func (c *Cache) DeleteByEmail(email string) error {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.items[email]; !ok {
		return fmt.Errorf("user %v was not found", email)
	}
	delete(c.items, email)
	var keys []string
	keys = append(keys, email)
	err := c.db.DeleteByEmail(keys)
	if err != nil {
		return err
	}
	// initializing map if it is nil
	c.updateMap()

	return nil
}

// controller for gc

func (c *Cache) garbageCollector() {
	<-time.After(c.cleanUpInterval)
	for {
		c.updateMap()
		// if expired items exist, delete them
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
			err := c.db.DeleteByEmail(keys)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

}

func (c *Cache) GetEmployeeInfoByEmail(email string) (*models.EmployeeInfo, error) {
	c.RLock()
	defer c.RUnlock()
	item, ok := c.items[email]
	if !ok {
		return nil, fmt.Errorf("user with email %v was not found", email)

	}
	if time.Now().Compare(item.Expiration) == 1 {
		return nil, fmt.Errorf("user with email %v is not available", email)
	}

	info, err := c.db.GetEmployeeInfoByEmail(email)

	return info, err

}

// finding expired items

func (c *Cache) expiredKeys() (keys []string) {
	c.RLock()
	defer c.RUnlock()
	// checking for expired items in cache
	for key, value := range c.items {

		if time.Now().Compare(value.Expiration) == 1 {
			keys = append(keys, key)

		}

	}

	return
}

// adding in cache and db

func (c *Cache) Set(info *models.EmployeeInfo) error {
	c.Lock()

	key := info.Email
	if _, ok := c.items[key]; ok {
		c.Unlock()
		return c.update(info)
	}

	err := c.db.InsertEmployee(*info)
	if err != nil {
		return err
	}
	c.items[key] = &models.Item{Created: time.Now(), Expiration: time.Now().Add(c.defaultExpiration)}
	c.Unlock()

	return nil
}

func (c *Cache) update(info *models.EmployeeInfo) error {
	c.Lock()
	defer c.Unlock()
	key := info.Email
	_, ok := c.items[key]
	if !ok {
		defer log.Fatal("internal error")
		return fmt.Errorf("internal error")
	}
	return c.db.UpdateEmployeeInfo(info)

}

// updating map if it's nil
func (c *Cache) updateMap() {
	if c.items == nil {
		c.items = make(map[string]*models.Item)
	}
}
