package routes

import (
	"github.com/ashishra0/raspberry-pi-service/httpd/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func StartService() {
	// PORT for prod environment
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	// To trigger the webhook
	router.GET("/api/power", handler.PostToSmee)
	// To save the response to db
	router.POST("/api/power/status", handler.HandlePost)
	// To return the saved response to user
	router.GET("/api/power/status", handler.GetStatus)
	// To handle the case of incorrect route
	router.NoRoute(func(context *gin.Context) {
		context.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":" + port)
}