package delivery

import (
	"SimpleServer/internal/controller"
	"SimpleServer/internal/delivery/v1"
	"SimpleServer/internal/providers/cache"
	"SimpleServer/internal/providers/db"
	"SimpleServer/internal/services/employmentService"
	"SimpleServer/pkg/userService"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "9340fk3__132AA@"
	dbName     = "company"
	driverName = "postgres"
)

func RunGRPCServer() error {

	// configuration for server
	time.Sleep(time.Second * 0)
	fmt.Println("Server is working")
	dataBase, err := db.NewDB(host, user, password, dbName, driverName, port)
	if err != nil {
		return err
	}
	currentCache := cache.NewCache(time.Second*30, time.Minute*0, dataBase)
	client, err := employmentService.Init("localhost:8082")
	if err != nil {
		log.Fatal(err.Error())
	}
	cnt := controller.NewController(currentCache)
	server := v1.NewGrpcServer(cnt, client)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	grpcSrv := grpc.NewServer()
	userService.RegisterUserServiceServer(grpcSrv, server)

	err = grpcSrv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
