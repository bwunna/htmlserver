package Cache_

import (
	"SimpleServer/Internal/App/Models"
	"fmt"
	"time"
)

func (c *Cache) Set(key string, value interface{}, duration time.Duration) error {
	var expiration time.Time
	var endlessLifeTime bool
	c.Lock()
	defer c.Unlock()
	if _, ok := c.items[key]; ok {
		return fmt.Errorf("user with name %v is not unique", key)
	}
	// checking for endless lifetime availability for item for this cache struct
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

	c.items[key] = Models.Item{Value: value, Created: time.Now(), Expiration: expiration, EndlessLifeTime: endlessLifeTime}
	fmt.Println(c.items)
	return nil
}
