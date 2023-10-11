package ServerStarter

import (
	"SimpleServer/Internal/App/Controllers/Controller"
	"SimpleServer/Internal/App/Providers/Provider"
	"time"
)

func StartServer() {
	Server := Controller.NewServer("localhost:8080", nil)
	db, err := Provider.NewDataBase("localhost", "postgres", "9340fk3__132AA@", "company", "postgres", 5432)
	if err != nil {
		panic(err.Error())
	}
	Server.Start(time.Minute, time.Minute*10, false, db, time.Minute*2)

}
