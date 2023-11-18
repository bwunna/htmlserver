package pkg

import (
	"SimpleServer/pkg/employmentService"
	"SimpleServer/pkg/userService"
)

func ConvertEmployeeInfoFromHeadHunterToUserService(info *employmentService.EmployeeInfo) *userService.Employee {

	return &userService.Employee{
		Name:           info.Name,
		Email:          info.Email,
		CompanyName:    info.CompanyName,
		DepartmentName: info.DepartmentName,
		Salary:         info.Salary,
		Status:         info.Status,
	}

}
