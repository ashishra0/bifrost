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
func StartService() {
		// PORT for prod environment
		port := os.Getenv("PORT")
		if port == "" {
				log.Fatal("$PORT must be set")
		}

		router := gin.Default()

		v1 := router.Group("/v1")
		{
				// To send a payload to smee.io
				v1.GET("/power", handler.TriggerWebhook)
				// To save the response to db
				v1.POST("/power/status", handler.HandlePost)
				// To return the saved response to user
				v1.GET("/power/status", handler.GetStatus)
				// To create a new expense record
				v1.POST("/expense/create", handler.CreateExpense)
				// To get the record from hasura
				v1.GET("/expense", handler.GetExpense)
				// To receive the event payload from hasura
				v1.POST("/expense", handler.FromHasura)
		}
		// To handle the case of incorrect route
		router.NoRoute(func(context *gin.Context) {
				context.AbortWithStatus(http.StatusNotFound)
		})
		router.Run(":" + port)
}
