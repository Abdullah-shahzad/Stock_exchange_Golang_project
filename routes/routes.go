package routes

import (
	"stock_exchange_Golang_project/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes() *gin.Engine {

	router := gin.Default()

	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/api/users/", controllers.CreateUser)
	router.GET("/api/users/:username/", controllers.GetUser)

	router.POST("/api/stocks/", controllers.CreateStock)
	router.GET("/api/stocks/", controllers.GetAllStocks)
	router.GET("/api/stocks/:ticker/", controllers.GetStockByTicker)

	router.POST("/api/transactions/", controllers.CreateTransaction)
	router.GET("/api/transactions/:username/", controllers.GetTransactions)
	router.GET("/api/transactions/:username/:start_time/:end_time/", controllers.GetTransactionsByDate)

	return router
}
