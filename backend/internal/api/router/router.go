package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karelpelcak/chat_call/internal/api/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	router := r.Group("/api/v1")
	{
		router.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		auth := router.Group("/auth")
		{
			auth.POST("/register", controller.HandleRegister)
		}
	}
	return r
}
