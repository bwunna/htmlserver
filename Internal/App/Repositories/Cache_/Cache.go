package Cache_

import (
	"SimpleServer/Internal/App/Models"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	items                       map[string]Models.Item
	defaultExpiration           time.Duration
	cleanUpInterval             time.Duration
	endlessLifeTimeAvailability bool
}
