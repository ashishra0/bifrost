package handler

import (
	"fmt"
	"github.com/ashishra0/raspberry-pi-service/model"
	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	message := model.MessageDB
	defaultMessage := "No power"
	if len(message) == 0 {
		fmt.Printf("status: %s", defaultMessage)
	} else {
		fmt.Printf("status: %s", model.MessageDB)
	}
}