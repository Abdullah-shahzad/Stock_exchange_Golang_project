package controllers

import (
	"database/sql"
	"net/http"
	"stock_exchange_Golang_project/config"

	"github.com/gin-gonic/gin"
)

type Stock struct {
	ID     int     `json:"id"`
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

func CreateStock(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	var stock Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO stocks (ticker, price) VALUES ($1, $2)`
	_, err := db.Exec(query, stock.Ticker, stock.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Stock created successfully"})
}

func GetAllStocks(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM stocks`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stocks"})
		return
	}
	defer rows.Close()

	var stocks []Stock
	for rows.Next() {
		var stock Stock
		err := rows.Scan(&stock.ID, &stock.Ticker, &stock.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse stock data"})
			return
		}
		stocks = append(stocks, stock)
	}

	c.JSON(http.StatusOK, stocks)
}

func GetStockByTicker(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	ticker := c.Param("ticker")

	var stock Stock
	query := `SELECT id, ticker, price FROM stocks WHERE ticker = $1`
	err := db.QueryRow(query, ticker).Scan(&stock.ID, &stock.Ticker, &stock.Price)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stock"})
		return
	}

	c.JSON(http.StatusOK, stock)
}
