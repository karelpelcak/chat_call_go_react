package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karelpelcak/chat_call/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type RequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func HnadleLogin(c *gin.Context) {}

func HandleRegister(c *gin.Context) {
	var body RequestBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	if body.Username == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	var existing string
	err := db.DB.Get(&existing, "SELECT username FROM users WHERE username=$1", body.Username)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already used"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	body.Password = string(hashedPassword)

	_, err = db.DB.NamedExec(`
		INSERT INTO users (username, password)
		VALUES (:username, :password)
	`, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}


func HandleMe(c *gin.Context) {}


