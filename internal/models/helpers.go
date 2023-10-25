package models

import "SimpleServer/pkg/usersService"

func ConvertGrpcUserToModelsUser(user *usersService.User) *User {
	return &User{Name: user.Name, Age: int(user.Age), Sex: user.Sex}
}
