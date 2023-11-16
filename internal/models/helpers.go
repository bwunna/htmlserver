package models

import (
	"SimpleServer/pkg/userService"
)

func ConvertToEmployeeResponse(info *EmployeeInfo) *userService.Employee {
	return &userService.Employee{
		Name:           info.Name,
		Email:          info.Email,
		DepartmentName: info.DepartmentName,
		CompanyName:    info.CompanyName,
		Salary:         info.Salary,
		Status:         info.Status,
	}
}

func ConvertToEmployeeInfo(info *userService.Employee) *EmployeeInfo {
	return &EmployeeInfo{
		Name:           info.Name,
		Email:          info.Email,
		DepartmentName: info.DepartmentName,
		CompanyName:    info.CompanyName,
		Salary:         info.Salary,
		Status:         info.Status,
	}
}
