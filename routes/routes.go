package routes

import (
	"stock_exchange_Golang_project/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes() *gin.Engine {

	router := gin.Default()

	router.POST("/users/", controllers.CreateUser)
	router.GET("/users/:username/", controllers.GetUser)

	router.POST("/stocks/", controllers.CreateStock)
	router.GET("/stocks/", controllers.GetAllStocks)
	router.GET("/stocks/:ticker/", controllers.GetStockByTicker)

	router.POST("/transactions/", controllers.CreateTransaction)
	router.GET("/transactions/:username/", controllers.GetTransactions)
	router.GET("/transactions/:username/:start_time/:end_time/", controllers.GetTransactionsByDate)

	return router
}

