package models

import "SimpleServer/pkg/usersService"

//func convertUserDeliveryTpUserDB(user *delivery.User) *db.User{
// return UserDB{}
//}

func ConvertUserFromGrpcPersonToModelsPerson(user *usersService.User) *User {
	return &User{Name: user.Name, Age: int(user.Age), Sex: user.Sex}
}
