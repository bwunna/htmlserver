package delivery

import (
	v1 "SimpleServer/internal/delivery/v1"
	"SimpleServer/internal/providers/cache"
	"SimpleServer/internal/providers/db"
	"SimpleServer/pkg/usersService"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

// В данном файле должна быть структура данного сервиса
// КОНСТРУКТОР
// фукнция запуска сервиса
//

const (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "9340fk3__132AA@"
	dbName     = "company"
	driverName = "postgres"
)

func RunGRPCServer() error {
	dataBase, err := db.NewDB(host, user, password, dbName, driverName, port)
	if err != nil {
		return fmt.Errorf("бд")
	}
	newCache := cache.NewCache(time.Minute*10, time.Minute*2, false, dataBase, time.Minute*2)
	server := v1.NewGrpcServer(newCache)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	grpcSrv := grpc.NewServer()
	usersService.RegisterUserCenterServer(grpcSrv, server)

	err = grpcSrv.Serve(lis)
	if err != nil {
		return fmt.Errorf("сёрв")
	}
	return nil
}

//func CreateAndRunUserCenter(address string, defaultExpiration time.Duration, cleanUpInterval time.Duration, endlessLifeTimeAvailability bool, db *Provider.DataBase, promotionInterval time.Duration) *v1.Server {
//	return nil
//
//}
