package controllers

import (
	"database/sql"
	"net/http"
	"stock_exchange_Golang_project/config"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

func CreateUser(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	var input struct {
		Username       string  `json:"username"`
		InitialBalance float64 `json:"initial_balance"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	query := `INSERT INTO users (username, balance) VALUES ($1, $2)`
	_, err := db.Exec(query, input.Username, input.InitialBalance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func GetUser(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	username := c.Param("username")

	var user User
	query := `SELECT id, username, balance FROM users WHERE username = $1`
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Balance)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
