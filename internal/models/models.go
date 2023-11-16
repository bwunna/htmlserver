package models

import "time"

// Structures for all project

type Item struct {
	Created    time.Time
	Expiration time.Time
}

type EmployeeInfo struct {
	Name, Email, CompanyName, DepartmentName, Status string
	Salary                                           int32
}
