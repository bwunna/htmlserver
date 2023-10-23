package main

import (
	"SimpleServer/internal/App/Services/ServerStarter"
	"log"
)

func main() {
	err := ServerStarter.StartServer()
	if err != nil {
		log.Fatal(err.Error())
	}

}
