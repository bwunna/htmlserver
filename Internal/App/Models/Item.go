package Models

import "time"

type Item struct {
	Value           interface{}
	Created         time.Time
	Expiration      time.Time
	EndlessLifeTime bool
}
