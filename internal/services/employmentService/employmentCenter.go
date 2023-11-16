package employmentService

import (
	"SimpleServer/pkg"
	"SimpleServer/pkg/employmentService"
	"SimpleServer/pkg/userService"
	"context"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	client employmentService.EmploymentCenterClient
}

// constructor for client sending requests to HH

func Init(url string) (*Client, error) {

	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error with connecting to head hunter %s", err.Error())
	}

	// Создание клиента второго сервиса
	return &Client{client: employmentService.NewEmploymentCenterClient(conn)}, nil

}

// getting employee info by his email

func (c *Client) GetEmployeeInfoByEmail(email string) (*userService.Employee, error) {
	ctx := context.Background()
	req := &employmentService.ByEmailRequest{Email: email}
	info, err := c.client.GetEmployeeInfoByEmail(ctx, req)

	if err != nil {
		return nil, err
	}
	return pkg.ConvertEmployeeInfoFromHeadHunterToUserService(info), nil
}
