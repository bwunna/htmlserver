package Cache_

import (
	"fmt"
	"time"
)

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
