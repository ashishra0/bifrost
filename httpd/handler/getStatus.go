package handler

import (
	"github.com/ashishra0/raspberry-pi-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func DefaultStatus() {
	model.MessageDB = "No Power"
}