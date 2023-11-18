package v1

import (
	"SimpleServer/internal/controller"
	"SimpleServer/internal/services/employmentService"
	"SimpleServer/pkg/userService"
	"context"
)

type Server struct {
	cnt    *controller.Controller
	client *employmentService.Client
}

// sending request to controller to add employee

func (s Server) AddEmployee(_ context.Context, employee *userService.Employee) (*userService.Basic, error) {
	err := s.cnt.AddEmployeeInCache(employee)
	if err != nil {
		return &userService.Basic{
			Code:    codeBAD,
			Message: err.Error(),
		}, err
	}
	return &userService.Basic{
		Code:    codeOK,
		Message: MessageOK,
	}, nil
}

// constructor for service api

func NewGrpcServer(cnt *controller.Controller, client *employmentService.Client) *Server {
	// constructor for api
	return &Server{
		cnt:    cnt,
		client: client,
	}
}

// sending request to controller to get employee info

func (s Server) GetEmployeeByEmail(_ context.Context, request *userService.EmailRequest) (*userService.Employee, error) {
	// checking if user exists
	info, err := s.cnt.GetEmployeeByEmail(request.Email)
	if err != nil {
		info, err = s.client.GetEmployeeInfoByEmail(request.Email)
		if err != nil {
			return nil, err
		}
		err = s.cnt.AddEmployeeInCache(info)
		if err != nil {
			return info, err
		}
	}

	return info, nil
}

// sending request to controller to delete the employee

func (s Server) DeleteEmployeeByEmail(_ context.Context, request *userService.EmailRequest) (*userService.Basic, error) {
	err := s.cnt.DeleteEmployeeByEmail(request.Email)
	// checking for successful deleting
	if err != nil {
		return &userService.Basic{Code: codeBAD, Message: err.Error()}, err
	}

	return &userService.Basic{
		Code:    codeOK,
		Message: MessageOK,
	}, nil
}
