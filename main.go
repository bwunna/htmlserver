package main

import (
	"SimpleServer/internal/delivery"
	"log"
)

func main() {
	err := delivery.RunGRPCServer()

	if err != nil {
		log.Fatal(err.Error())
	}

}
