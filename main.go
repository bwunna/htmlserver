package main

import (
	"SimpleServer/Internal/App/Services/ServerStarter"
	"log"
)

func main() {
	err := ServerStarter.StartServer()
	if err != nil {
		log.Fatal(err.Error())
	}

}
