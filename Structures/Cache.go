package Structures

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	mutex                       sync.RWMutex
	items                       map[string]Item
	defaultExpiration           time.Duration
	cleanUpInterval             time.Duration
	endlessLifeTimeAvailability bool
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	item, ok := c.items[key]
	if !ok {
		c.mutex.RUnlock()
		return nil, false
	}
	if !item.endlessLifeTime && time.Now().Compare(item.Expiration) == -1 {
		return nil, false
	}
	c.mutex.RUnlock()
	return item, true

}

func (c *Cache) Delete(key string) error {
	c.mutex.Lock()
	if _, ok := c.items[key]; !ok {
		c.mutex.Unlock()
		return errors.New("couldn't find the user")
	}
	delete(c.items, key)
	c.mutex.Unlock()
	return nil
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) error {
	var expiration time.Time
	var endlessLifeTime bool
	c.mutex.Lock()
	if _, ok := c.items[key]; ok {
		c.mutex.Unlock()
		return errors.New("key is not unique")

	}
	if duration == 0 {
		if c.endlessLifeTimeAvailability {
			endlessLifeTime = true
		} else {
			duration = c.defaultExpiration
		}
	}
	if duration > 0 {
		expiration = time.Now().Add(duration)
	}

	c.items[key] = Item{value, time.Now(), expiration, endlessLifeTime}
	c.mutex.Unlock()
	return nil
}
func (c *Cache) Update(key string, age int, sex bool) error {
	c.mutex.Lock()
	value, ok := c.items[key]
	if !ok {
		c.mutex.Unlock()
		return errors.New("couldn't find the user")

	} else {
		if person, ok := value.Value.(Person); ok {
			person.Age = age
			person.Sex = sex

		}
	}
	c.mutex.Unlock()
	return nil
}
func (c *Cache) gc() {
	<-time.After(c.cleanUpInterval)
	for {
		if c.items == nil {
			return
		}
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}

}
func (c *Cache) expiredKeys() (keys []string) {
	c.mutex.RLock()
	for key, value := range c.items {
		if !value.endlessLifeTime && time.Now().Compare(value.Expiration) == -1 {
			keys = append(keys, key)

		}

	}
	c.mutex.RUnlock()
	return
}
func (c *Cache) clearItems(keys []string) {
	c.mutex.Lock()
	for _, key := range keys {
		delete(c.items, key)
	}
	c.mutex.Unlock()
}

func NewCache(defaultExpiration, cleanupInterval time.Duration, endlessLifeTimeAvailability bool) *Cache {

	// инициализируем карту(map) в паре ключ(string)/значение(Item)
	items := make(map[string]Item)
	cache := Cache{

		items:                       items,
		defaultExpiration:           defaultExpiration,
		cleanUpInterval:             cleanupInterval,
		endlessLifeTimeAvailability: endlessLifeTimeAvailability,
	}

	// Если интервал очистки больше 0, запускаем GC (удаление устаревших элементов)
	go cache.gc()

	return &cache
}
