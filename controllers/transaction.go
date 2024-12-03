//controller.transaction.go
package controllers

import (
	"database/sql"
	"net/http"
	"stock_exchange_Golang_project/config"
	"time"

	"github.com/gin-gonic/gin"
)


type Transaction struct {
	ID                int     `json:"id"`
	UserID            int     `json:"user_id"`
	Ticker            string  `json:"ticker"`
	TransactionType   string  `json:"transaction_type"`
	TransactionVolume int     `json:"transaction_volume"`
	TransactionPrice  float64 `json:"transaction_price"`
	Timestamp         string  `json:"timestamp"`
}


type TransactionRequest struct {
	Username          string `json:"username" example:"abdullah"`
	Ticker            string `json:"ticker" example:"tia"`
	TransactionType   string `json:"transaction_type" example:"BUY"`
	TransactionVolume int    `json:"transaction_volume" example:"10"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Description Creates a new transaction record in the system.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param transaction body controllers.TransactionRequest true "Transaction data"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions [post]
func CreateTransaction(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	var input TransactionRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	if input.TransactionType != "BUY" && input.TransactionType != "SELL" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "TransactionType must be either BUY or SELL"})
		return
	}

	var userBalance float64
	err := db.QueryRow(`SELECT balance FROM users WHERE username = $1`, input.Username).Scan(&userBalance)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching user balance"})
		return
	}

	var stockPrice float64
	err = db.QueryRow(`SELECT price FROM stocks WHERE ticker = $1`, input.Ticker).Scan(&stockPrice)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Stock not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching stock price"})
		return
	}

	transactionPrice := stockPrice * float64(input.TransactionVolume)

	
	if input.TransactionType == "BUY" && userBalance < transactionPrice {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Insufficient balance"})
		return
	}

	
	var balanceQuery string
	if input.TransactionType == "BUY" {
		balanceQuery = `UPDATE users SET balance = balance - $1 WHERE username = $2`
	} else {
		balanceQuery = `UPDATE users SET balance = balance + $1 WHERE username = $2`
	}
	_, err = db.Exec(balanceQuery, transactionPrice, input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to update balance"})
		return
	}

	// Insert transaction into the database
	insertQuery := `
        INSERT INTO transactions (user_id, ticker, transaction_type, transaction_volume, transaction_price, timestamp)
        VALUES ((SELECT id FROM users WHERE username = $1), $2, $3, $4, $5, $6)`
	_, err = db.Exec(insertQuery, input.Username, input.Ticker, input.TransactionType, input.TransactionVolume, transactionPrice, time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create transaction"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, SuccessResponse{Message: "Transaction completed successfully"})
}

// GetTransactions godoc
// @Summary Get all Transactions for a user
// @Description Retrieves a list of all transactions for a given user.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {array} Transaction
// @Failure 500 {object} ErrorResponse
// @Router /transactions/{username} [get]
func GetTransactions(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	username := c.Param("username")

	query := `
		SELECT t.id, t.ticker, t.transaction_type, t.transaction_volume, t.transaction_price, t.timestamp
		FROM transactions t
		INNER JOIN users u ON t.user_id = u.id
		WHERE u.username = $1
		ORDER BY t.timestamp DESC`

	rows, err := db.Query(query, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.ID, &transaction.Ticker, &transaction.TransactionType, &transaction.TransactionVolume, &transaction.TransactionPrice, &transaction.Timestamp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing transactions"})
			return
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// GetTransactionsByDate godoc
// @Summary Get Transactions for a user by timestamp
// @Description Retrieves transactions for a specific user within a given time range.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param username path string true "Username of the user"
// @Param start_time path string true "Start timestamp in YYYY-MM-DD format" format(date)
// @Param end_time path string true "End timestamp in YYYY-MM-DD format" format(date)
// @Success 200 {array} Transaction
// @Failure 500 {object} ErrorResponse
// @Router /transactions/{username}/{start_time}/{end_time} [get]
func GetTransactionsByDate(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	username := c.Param("username")
	startTime := c.Param("start_time")
	endTime := c.Param("end_time")

	query := `
		SELECT t.id, t.ticker, t.transaction_type, t.transaction_volume, t.transaction_price, t.timestamp
		FROM transactions t
		INNER JOIN users u ON t.user_id = u.id
		WHERE u.username = $1 AND t.timestamp BETWEEN $2 AND $3
		ORDER BY t.timestamp DESC`

	rows, err := db.Query(query, username, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.ID, &transaction.Ticker, &transaction.TransactionType, &transaction.TransactionVolume, &transaction.TransactionPrice, &transaction.Timestamp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing transactions"})
			return
		}
		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
