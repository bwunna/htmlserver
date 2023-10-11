package Cache_

import (
	"SimpleServer/Internal/App/Models"
	"SimpleServer/Internal/App/Providers/Provider"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	items                       map[string]Models.Item
	defaultExpiration           time.Duration
	cleanUpInterval             time.Duration
	endlessLifeTimeAvailability bool
	promotionInterval           time.Duration
	db                          *Provider.DataBase
}

func (c *Cache) CheckForItem(key string) bool {
	_, ok := c.items[key]

	return ok
}

func (c *Cache) GetSalaryData(key string) (*Models.TableData, error) {
	// returns data about employee salary
	employeeInfo, err := c.db.GetEmployeeInfo(key)
	if err != nil {
		return nil, err
	} else {
		return employeeInfo, nil
	}
}
func (c *Cache) clearItems(keys []string) {
	c.Lock()

	defer c.Unlock()
	// clearing items by their keys
	for _, key := range keys {
		delete(c.items, key)
	}

}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()
	// return error if item was not found
	if _, ok := c.items[key]; !ok {
		return fmt.Errorf("user %v was not found", key)
	}
	delete(c.items, key)
	var keys []string
	keys = append(keys, key)
	err := c.db.Delete(keys)
	if err != nil {
		return err
	}
	// initializing map if it is nil
	c.updateMap()

	return nil
}

func (c *Cache) garbageCollector() {
	<-time.After(c.cleanUpInterval)
	for {
		// initializing map if it is nil
		c.updateMap()
		// if expired items exist, delete them
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
			err := c.db.Delete(keys)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

}

func (c *Cache) Get(key string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()
	item, ok := c.items[key]
	//  if item was not found
	if !ok {
		return nil, fmt.Errorf("user with key %v was not found", key)

	}
	//  if item is expired
	if !item.EndlessLifeTime && time.Now().Compare(item.Expiration) == 1 {
		return nil, fmt.Errorf("user with key %v is not available", key)
	}

	return item, nil

}

func (c *Cache) AskForPromotion(key string) error {
	if value, ok := c.items[key]; !ok {
		return fmt.Errorf("error: user not found")
	} else {
		if value.Created.Add(c.promotionInterval).Compare(time.Now()) == 1 {

			return fmt.Errorf("error: you need to work more")
		} else {
			err := c.db.AskForPromotion(key)
			return err
		}
	}

}

func (c *Cache) expiredKeys() (keys []string) {
	c.RLock()
	defer c.RUnlock()
	// checking for expired items in cache
	for key, value := range c.items {

		if !value.EndlessLifeTime && time.Now().Compare(value.Expiration) == 1 {
			keys = append(keys, key)

		}

	}

	return
}
func NewCache(defaultExpiration, cleanupInterval time.Duration, endlessLifeTimeAvailability bool, db *Provider.DataBase, promotionInterval time.Duration) *Cache {

	// initializing map
	items := make(map[string]Models.Item)
	cache := Cache{

		items:                       items,
		defaultExpiration:           defaultExpiration,
		cleanUpInterval:             cleanupInterval,
		endlessLifeTimeAvailability: endlessLifeTimeAvailability,
		db:                          db,
		promotionInterval:           promotionInterval,
	}

	// starting gc
	go cache.garbageCollector()

	return &cache
}

func (c *Cache) ParseJson(decoder *json.Decoder) (*Models.Person, error) {
	var person Models.Person
	err := decoder.Decode(&person)
	fmt.Println(person)
	if err != nil {
		return nil, fmt.Errorf("error: invalid json file")
	}

	return &person, nil
}

func (c *Cache) Set(person *Models.Person, duration time.Duration) error {
	var expiration time.Time
	var endlessLifeTime bool

	c.Lock()
	defer c.Unlock()
	key := person.Name
	if _, ok := c.items[key]; ok {
		return fmt.Errorf("user with name %v is not unique", key)
	}
	// checking for endless lifetime availability for item from this cache
	if duration == 0 {
		if c.endlessLifeTimeAvailability {
			endlessLifeTime = true
		} else {
			duration = c.defaultExpiration
		}
	}
	// counting expiration for this item
	if duration > 0 {
		expiration = time.Now().Add(duration)
	}
	err := c.db.Insert(person.Name)
	if err != nil {
		panic(err.Error())
	}
	c.items[key] = Models.Item{Value: person, Created: time.Now(), Expiration: expiration, EndlessLifeTime: endlessLifeTime}

	return nil
}

func (c *Cache) Update(person *Models.Person) error {
	c.Lock()
	defer c.Unlock()
	key := person.Name
	value, ok := c.items[key]
	// updates info about user
	if !ok {
		c.Unlock()
		return errors.New("couldn't find the user")
		// if user was not found, return error
	} else {
		if user, ok := value.Value.(*Models.Person); ok {
			user.Age = person.Age
			user.Sex = person.Sex
		} else {
			return fmt.Errorf("couldn't post this type")
		}
	}

	return nil
}

// updating map if it's nil
func (c *Cache) updateMap() {
	if c.items == nil {
		c.items = make(map[string]Models.Item)
	}
}
