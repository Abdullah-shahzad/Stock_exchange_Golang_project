//controller.user.go

package controllers

import (
	"database/sql"
	"net/http"
	"stock_exchange_Golang_project/config"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}

type UserRequest struct {
	Username       string  `json:"username" example:"abdullah"`
	InitialBalance float64 `json:"initial_balance" example:"1000.00"`
}


// CreateUser godoc
// @Summary Create a new user
// @Description Saves new user data into the database.
// @Tags User
// @Accept json
// @Produce json
// @Param user body controllers.UserRequest true "User data"
// @Success 201 {object} controllers.SuccessResponse
// @Failure 400 {object} controllers.ErrorResponse
// @Failure 500 {object} controllers.ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	var input UserRequest
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


// GetUser godoc
// @Summary get user by username
// @Description Retrieves user details based on the provided username.
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Success 200 {object} User
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{username} [get]
func GetUser(c *gin.Context) {
	db := config.ConnectDB()
	defer db.Close()

	username := strings.TrimSpace(c.Param("username"))

	var user User
	query := `SELECT id, username, balance FROM users WHERE LOWER(username) = LOWER($1)`
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
