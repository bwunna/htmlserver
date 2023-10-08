package Cache_

import (
	"SimpleServer/Internal/App/Models"
	"fmt"
)

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()
	// return error if item was not found
	if _, ok := c.items[key]; !ok {
		return fmt.Errorf("user %v was not found", key)
	}
	delete(c.items, key)
	// initializing map if it is nil
	if c.items == nil {
		c.items = make(map[string]Models.Item)
	}

	return nil
}
