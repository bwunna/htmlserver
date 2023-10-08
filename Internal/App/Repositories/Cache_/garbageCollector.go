package Cache_

import (
	"SimpleServer/Internal/App/Models"
	"time"
)

func (c *Cache) garbageCollector() {
	<-time.After(c.cleanUpInterval)
	for {
		// initializing map if it is nil
		if c.items == nil {
			c.items = make(map[string]Models.Item)
		}
		// if expired items exist, delete them
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}

}
