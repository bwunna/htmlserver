package v1

import (
	"SimpleServer/internal/models"
	"SimpleServer/internal/providers/cache"
	"SimpleServer/pkg/usersService"
	"context"
	"fmt"
	"time"
)

// набор методов апи в сервере

// NOT USE PROVIDERS PACKAGE HERE. ONLY controller func ans structs

// func (method)createUser(){
//	controller.CreateUser()
//  controller.UpdateStatus
//}

//func GetUserByName(request *grpcserver.GetUserByNameRequest) (*grpcserver.BasicResponse, error){
// if err != nil{
// far far away logger.Info("error in GetUserByName func" , err)

// return &grpcserver.BasicResponse{
//code: errorCode} , nil
//}

// return &grpcserver.BasicResponse{} , nil

// }
//type Server struct {
//	address string
//	cache   *cache.Cache
//}
//*cache
//func NewServer(address string) *Server {
//	// new server
//	return &Server{address: address}
//}

type GrpcServer struct {
	cache *cache.Cache
}

func NewGrpcServer(cache *cache.Cache) *GrpcServer {
	return &GrpcServer{cache: cache}
}

func (s GrpcServer) UpdateSalary(_ context.Context, request *usersService.UserByNameRequest) (*usersService.BasicResponse, error) {

	err := s.cache.AskForPromotion(request.Name)
	if err != nil {
		return new(usersService.BasicResponse), err
	}
	return new(usersService.BasicResponse), nil
	/*_, err = fmt.Fp3rintln(writer, "Congratulations! You've been promoted")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusConflict)
	}*/
}

func (s GrpcServer) GetUserByName(_ context.Context, request *usersService.UserByNameRequest) (*usersService.EmployeeInfo, error) {
	// checking if user exists
	empInfo := &usersService.EmployeeInfo{Info: ""}
	userData, err := s.cache.Get(request.Name)
	if err != nil {
		return empInfo, err
	}

	salaryData, err := s.cache.GetSalaryData(request.Name)
	if err != nil {
		return empInfo, err
	}
	empInfo.Info = fmt.Sprint(*userData, salaryData)
	return empInfo, nil
}

func (s GrpcServer) DeleteUserByName(_ context.Context, request *usersService.UserByNameRequest) (*usersService.BasicResponse, error) {
	// checking for valid key
	err := s.cache.Delete(request.Name)
	// checking for successful deleting
	basicResponse := &usersService.BasicResponse{}
	if err != nil {
		return basicResponse, err
	}
	return basicResponse, nil

}

func (s GrpcServer) AddUser(_ context.Context, user *usersService.User) (*usersService.BasicResponse, error) {
	// checking for successful decoding person from json
	//fmt.Println(request.Body)
	validUser := models.ConvertUserFromGrpcPersonToModelsPerson(user)
	basicResponse := &usersService.BasicResponse{}
	err := s.cache.Set(validUser, time.Minute*10)
	//fmt.Println(request.Body)
	if err != nil {
		// checking for errors while adding the user
		if err = s.cache.Update(validUser); err != nil {
			return basicResponse, err
		}
	}
	return basicResponse, nil

}
