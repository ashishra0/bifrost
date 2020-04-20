package routes

import (
	"github.com/ashishra0/raspberry-pi-service/httpd/handler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

// StartService runs the gin application.
// It is the single point of entry in the application.
func StartTeslaService() {
	// PORT for prod environment
	port := "5000"
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	// To trigger the webhook
	router.GET("/api/power", handler.TriggerWebhook)
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


func StartAlfredService() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	// To get the data from hasura
	router.GET("/api/expense", handler.GetExpense)
	// To forward the event response from hasura
	router.POST("/api/expense", handler.PostExpense)
	router.NoRoute(func(context *gin.Context) {
		context.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":" + port)
}