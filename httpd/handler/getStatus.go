package handler

import (
	"github.com/ashishra0/raspberry-pi-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStatus will retrieve the saved status string from messageDB defined in model/message.go
// It sends a response of the saved status or a default status if not present.
func GetStatus(c *gin.Context) {
	message := model.MessageDB
	defaultMessage := "No Power"
	if len(message) == 0 {
		c.String(http.StatusOK, defaultMessage)
	} else {
		c.String(http.StatusOK, model.MessageDB)
	}
	DefaultStatus()
}

// DefaultStatus sets the messageDB variable to a default message
func DefaultStatus() {
	model.MessageDB = "No Power"
}