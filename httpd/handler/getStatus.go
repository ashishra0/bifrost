package handler

import (
	"github.com/ashishra0/raspberry-pi-service/helper"
	"github.com/ashishra0/raspberry-pi-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStatus will retrieve the saved status string from messageDB defined in model/message.go
// It sends a response of the saved status or a default status if not present.
func GetStatus(c *gin.Context) {
	message := model.MessageDB
	errorMessage := helper.GetError()
	if len(message) == 0 {
		c.JSON(http.StatusOK, errorMessage)
	} else {
		c.JSON(http.StatusOK, model.MessageDB)
	}
	helper.SetDefaultStatus()
}
