package main

import (
	"log"
	_ "stock_exchange_Golang_project/docs"
	"stock_exchange_Golang_project/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Stock Exchange API
// @description This is the API documentation for the Stock Exchange project
// @contact.name abdullah
// @contact.email abdullahkpr22@gmail.com
// @securityDefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /api
func main() {

	router := routes.ConfigureRoutes()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Server is running on port 8080...")
	log.Fatal(router.Run(":8080"))
}
