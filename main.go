package main

import (
	"log"
	_ "stock_exchange_Golang_project/docs" // Import docs for swagger
	"stock_exchange_Golang_project/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Stock exchange API
// @description A stock exchange API project using the Gin framework

// @host localhost:8080
// @BasePath /api
func main() {

	router := routes.ConfigureRoutes()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server is running on port 8080...")
	log.Fatal(router.Run(":8080"))
}
