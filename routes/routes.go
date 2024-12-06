package routes

import (
	"stock_exchange_Golang_project/controllers"
	"stock_exchange_Golang_project/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes() *gin.Engine {

	router := gin.Default()

	router.Use(middleware.DBMiddleware())

	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutes := router.Group("/user")
	{
		authRoutes.POST("/register", controllers.Signup)
		authRoutes.POST("/login", controllers.Login)
		authRoutes.GET("/authenticated", middleware.AuthMiddleware, controllers.IsAuthenticated)
	}

	userRoutes := router.Group("/api/users")
	{
		userRoutes.POST("/", middleware.AuthMiddleware, controllers.CreateUser)
		userRoutes.GET("/:username/", middleware.AuthMiddleware, controllers.GetUser)
	}

	stockRoutes := router.Group("/api/stocks")
	{
		stockRoutes.POST("/", middleware.AuthMiddleware, controllers.CreateStock)
		stockRoutes.GET("/", controllers.GetAllStocks)
		stockRoutes.GET("/:ticker", middleware.AuthMiddleware, controllers.GetStockByTicker)
	}

	transactionRoutes := router.Group("/api/transactions")
	{
		transactionRoutes.POST("/", middleware.AuthMiddleware, controllers.CreateTransaction)
		transactionRoutes.GET("/:username/", middleware.AuthMiddleware, controllers.GetTransactions)
		transactionRoutes.GET("/:username/:start_time/:end_time/", middleware.AuthMiddleware, controllers.GetTransactionsByDate)
	}

	return router
}
