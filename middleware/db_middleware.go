package middleware

import (
	"stock_exchange_Golang_project/config"
	"github.com/gin-gonic/gin"
)

func DBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		db := config.ConnectDB()

		c.Set("db", db)

		defer db.Close()

		c.Next()
	}
}
