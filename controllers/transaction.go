package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"stock_exchange_Golang_project/config"
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

func CreateTransaction(c *gin.Context) {
	
	db := config.ConnectDB()
	defer db.Close()

	var input struct {
		Username          string `json:"username"`
		Ticker            string `json:"ticker"`
		TransactionType   string `json:"transaction_type"`
		TransactionVolume int    `json:"transaction_volume"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if input.TransactionVolume <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaction volume must be greater than zero"})
		return
	}

	if input.TransactionType != "BUY" && input.TransactionType != "SELL" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction type"})
		return
	}

	var userBalance float64
	err := db.QueryRow(`SELECT balance FROM users WHERE username = $1`, input.Username).Scan(&userBalance)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user balance"})
		return
	}

	var stockPrice float64
	err = db.QueryRow(`SELECT price FROM stocks WHERE ticker = $1`, input.Ticker).Scan(&stockPrice)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching stock price"})
		return
	}

	transactionPrice := stockPrice * float64(input.TransactionVolume)

	if input.TransactionType == "BUY" && userBalance < transactionPrice {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	balanceQuery := ``
	if input.TransactionType == "BUY" {
		balanceQuery = `UPDATE users SET balance = balance - $1 WHERE username = $2`
	} else {
		balanceQuery = `UPDATE users SET balance = balance + $1 WHERE username = $2`
	}
	_, err = db.Exec(balanceQuery, transactionPrice, input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update balance"})
		return
	}

	insertQuery := `
		INSERT INTO transactions (user_id, ticker, transaction_type, transaction_volume, transaction_price, timestamp)
		VALUES ((SELECT id FROM users WHERE username = $1), $2, $3, $4, $5, $6)`
	_, err = db.Exec(insertQuery, input.Username, input.Ticker, input.TransactionType, input.TransactionVolume, transactionPrice, time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction completed successfully"})
}

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
