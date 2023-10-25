package delivery

import (
	"SimpleServer/internal/controller"
	v1 "SimpleServer/internal/delivery/v1"
	"SimpleServer/internal/providers/cache"
	"SimpleServer/internal/providers/db"
	"SimpleServer/pkg/usersService"
	"google.golang.org/grpc"
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
	dataBase, err := db.NewDB(host, user, password, dbName, driverName, port)
	if err != nil {
		return err
	}
	newCache := cache.NewCache(time.Minute*10, time.Minute*2, false, dataBase, time.Minute*2)
	newController := controller.NewController(newCache)
	server := v1.NewGrpcServer(newController)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	grpcSrv := grpc.NewServer()
	usersService.RegisterUserCenterServer(grpcSrv, server)

	err = grpcSrv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
