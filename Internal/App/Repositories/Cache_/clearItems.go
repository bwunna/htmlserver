package Cache_

func (c *Cache) clearItems(keys []string) {
	c.Lock()
	defer c.Unlock()
	// clearing items by their keys
	for _, key := range keys {
		delete(c.items, key)
	}

}
