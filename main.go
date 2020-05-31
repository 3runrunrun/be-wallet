package main

import (
	"log"

	"github.com/3runrunrun/be-wallet/server"
)

func main() {
	serve := server.Server()
	log.Println("be-wallet running on PORT 3681")

	err := serve.Run(":3681")
	if err != nil {
		log.Panic("be-wallet error...")
	}
}
