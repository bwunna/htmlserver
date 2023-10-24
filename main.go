package main

import (
	"SimpleServer/internal/app/services/serverStarter"
	"log"
)

func main() {
	err := serverStarter.StartServer()
	if err != nil {
		log.Fatal(err.Error())
	}

}
