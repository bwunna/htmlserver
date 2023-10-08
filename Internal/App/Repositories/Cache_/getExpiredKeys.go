package Cache_

import (
	"time"
)

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
