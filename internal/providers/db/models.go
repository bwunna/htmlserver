package db

import "time"

type TableData struct {
	Name           string
	EmploymentTime time.Time
	Salary         int
	Status         string
}
