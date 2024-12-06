package controllers

import (
	"database/sql"
	"net/http"
	"stock_exchange_Golang_project/models"
	"stock_exchange_Golang_project/utils/auth"

	"github.com/gin-gonic/gin"
)

type SignupResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Signup godoc
// @Summary Register auth-user
// @Description Add a new user
// @Tags auth_user
// @Accept json
// @Produce json
// @Param user body models.A_user true "A_user Data"
// @Success 201 {object} SignupResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/register [post]
func Signup(c *gin.Context) {
	var user models.A_user

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid input"})
		return
	}

	if err := user.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error hashing password"})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	query := `INSERT INTO auth_user (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Error creating stock",
		})
		return
	}

	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error generating token"})
		return
	}

	c.JSON(http.StatusCreated, SignupResponse{
		Message: "User created successfully",
		Token:   token,
	})
}

// Login godoc
// @Summary Login user
// @Description Login with username and password
// @Tags auth_user
// @Accept json
// @Produce json
// @Param creds body controllers.LoginCredentials true "Login credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/login [post]
func Login(c *gin.Context) {
	var creds LoginCredentials

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid input",
		})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	var user models.A_user
	query := `SELECT id, username, email, password FROM auth_user WHERE username = $1`
	err := db.QueryRow(query, creds.Username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid credentials"})
		return
	}

	if !user.CheckPassword(creds.Password) {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid credentials"})
		return
	}

	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
	})
}

// IsAuthenticated godoc
// @Summary Check if user is authenticated
// @Description Validate the JWT token
// @Tags auth_user
// @Accept json
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 401 {object} ErrorResponse
// @Router /user/authenticated [get]
func IsAuthenticated(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "No token provided"})
		c.Abort()
		return
	}

	claims, err := auth.ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid token"})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
	c.Next()
}
