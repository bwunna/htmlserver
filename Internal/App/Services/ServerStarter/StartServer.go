package ServerStarter

import (
	"SimpleServer/Internal/App/Controllers/ServerAPI"
	"time"
)

func StartServer() {
	Server := ServerAPI.New("localhost:8080", nil)
	Server.Start(time.Minute, time.Minute*2, false)
}
