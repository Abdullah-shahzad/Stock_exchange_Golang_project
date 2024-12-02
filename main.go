package main

import (
	"log"
	"stock_exchange_Golang_project/routes"
)

func main() {

	router := routes.ConfigureRoutes()

	log.Println("Server is running on port 8080...")
	log.Fatal(router.Run(":8080"))
}
