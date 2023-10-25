package models

import "time"

// Structures for all project

type Item struct {
	Value           interface{}
	Created         time.Time
	Expiration      time.Time
	EndlessLifeTime bool
}

type User struct {
	Name string
	Age  int
	Sex  bool
}
