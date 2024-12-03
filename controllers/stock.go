//controller.stock.go
package controllers

import (
	"database/sql"
	"net/http"
	"stock_exchange_Golang_project/config"

	"github.com/gin-gonic/gin"
)

// Stock represents a stock entity in the database.
type Stock struct {
	ID     int     `json:"id"`
	Ticker string  `json:"ticker" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

type CreateStockRequest struct {
	Ticker string  `json:"ticker" example:"AAPL"`
	Price  float64 `json:"price" example:"150.25"`
}

// CreateStock godoc
// @Summary Create a new stock entry
// @Description Saves new stock data into the database.
// @Tags Stock
// @Accept json
// @Produce json
// @Param stock body CreateStockRequest true "Stock data"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /stocks [post]
func CreateStock(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	var stock CreateStockRequest
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input. Ensure 'ticker' and 'price' are provided."})
		return
	}

	query := `INSERT INTO stocks (ticker, price) VALUES ($1, $2)`
	_, err := db.Exec(query, stock.Ticker, stock.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock in the database."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Stock created successfully."})
}

// GetAllStocks godoc
// @Summary Retrieve all stocks
// @Description Retrieves all stock entries from the database.
// @Tags Stock
// @Accept json
// @Produce json
// @Success 200 {array} Stock
// @Failure 500 {object} ErrorResponse
// @Router /stocks [get]
func GetAllStocks(c *gin.Context) {
	// Connect to the database
	db := config.ConnectDB()
	defer db.Close()

	// Query to fetch all stocks
	rows, err := db.Query(`SELECT id, ticker, price FROM stocks`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to retrieve stocks from the database.",
		})
		return
	}
	defer rows.Close()

	// Parse the rows into a slice of Stock
	var stocks []Stock
	for rows.Next() {
		var stock Stock
		err := rows.Scan(&stock.ID, &stock.Ticker, &stock.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error: "Failed to parse stock data.",
			})
			return
		}
		stocks = append(stocks, stock)
	}

	// Return the list of stocks
	c.JSON(http.StatusOK, stocks)
}

// GetStockByTicker godoc
// @Summary Retrieve stock by ticker
// @Description Retrieves stock details based on the provided ticker symbol.
// @Tags Stock
// @Accept json
// @Produce json
// @Param ticker path string true "Stock Ticker"
// @Success 200 {object} Stock
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /stocks/{ticker} [get]
func GetStockByTicker(c *gin.Context) {
	// Connect to the database
	db := config.ConnectDB()
	defer db.Close()

	// Retrieve ticker from the URL parameter
	ticker := c.Param("ticker")

	// Query to fetch stock by ticker
	var stock Stock
	query := `SELECT id, ticker, price FROM stocks WHERE ticker = $1`
	err := db.QueryRow(query, ticker).Scan(&stock.ID, &stock.Ticker, &stock.Price)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error: "Stock not found.",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to retrieve stock.",
		})
		return
	}

	// Return the stock data
	c.JSON(http.StatusOK, stock)
}
