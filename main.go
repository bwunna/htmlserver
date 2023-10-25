package main

import (
	"SimpleServer/internal/delivery"
	"fmt"
	"log"
)

func main() {
	fmt.Println("server is working")
	err := delivery.RunGRPCServer()

	if err != nil {
		log.Fatal(err.Error())
	}

}
