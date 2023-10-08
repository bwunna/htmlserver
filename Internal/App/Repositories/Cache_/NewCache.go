package Cache_

import (
	"SimpleServer/Internal/App/Models"
	"time"
)

func NewCache(defaultExpiration, cleanupInterval time.Duration, endlessLifeTimeAvailability bool) *Cache {

	// initializing map
	items := make(map[string]Models.Item)
	cache := Cache{

		items:                       items,
		defaultExpiration:           defaultExpiration,
		cleanUpInterval:             cleanupInterval,
		endlessLifeTimeAvailability: endlessLifeTimeAvailability,
	}

	// starting gc
	go cache.garbageCollector()

	return &cache
}
