package routes

import (
	"github.com/ashishra0/raspberry-pi-service/httpd/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func StartService() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	router.GET("/", handler.PostToSmee)
	router.POST("/", handler.HandlePost)
	router.GET("/api/power/status", handler.GetStatus)
	router.NoRoute(func(context *gin.Context) {
		context.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":" + port)
}