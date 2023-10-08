package ServerStarter

import (
	"SimpleServer/Internal/App/Controllers/Controller"
	"time"
)

func StartServer() {
	Server := Controller.NewServer("localhost:8080", nil)
	Server.Start(time.Minute, time.Minute*2, false)
}
