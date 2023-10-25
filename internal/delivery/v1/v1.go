package v1

import (
	"SimpleServer/internal/controller"
	"SimpleServer/internal/models"
	"SimpleServer/pkg/usersService"
	"context"
)

type GrpcServer struct {
	cnt *controller.Controller
}

func NewGrpcServer(cnt *controller.Controller) *GrpcServer {
	// constructor for api
	return &GrpcServer{cnt: cnt}
}

func (s GrpcServer) UpdateSalary(_ context.Context, request *usersService.UserByNameRequest) (*usersService.BasicResponse, error) {
	// updating salary
	err := s.cnt.UpdateSalary(request.Name)
	response := new(usersService.BasicResponse)
	if err != nil {
		return new(usersService.BasicResponse), err
	}
	response.Message = "Congratulations!"

	return new(usersService.BasicResponse), nil
}

func (s GrpcServer) GetUserByName(_ context.Context, request *usersService.UserByNameRequest) (*usersService.EmployeeInfo, error) {
	// checking if user exists
	empInfo := &usersService.EmployeeInfo{Info: ""}
	userData, err := s.cnt.GetUser(request.Name)
	if err != nil {
		return empInfo, err
	}

	empInfo.Info = userData

	return empInfo, nil
}

func (s GrpcServer) DeleteUserByName(_ context.Context, request *usersService.UserByNameRequest) (*usersService.BasicResponse, error) {
	err := s.cnt.DeleteUser(request.Name)
	// checking for successful deleting
	basicResponse := &usersService.BasicResponse{}
	if err != nil {
		return basicResponse, err
	}
	basicResponse.Message = "successful"

	return basicResponse, nil

}

func (s GrpcServer) AddUser(_ context.Context, user *usersService.User) (*usersService.BasicResponse, error) {
	// adding user
	basicResponse := &usersService.BasicResponse{}
	err := s.cnt.AddUser(models.ConvertGrpcUserToModelsUser(user))
	if err != nil {
		return basicResponse, nil
	}
	basicResponse.Message = "successful"
	return basicResponse, nil

}
