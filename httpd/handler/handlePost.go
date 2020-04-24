package handler

import (
				"github.com/ashishra0/raspberry-pi-service/model"
				"github.com/gin-gonic/gin"
				"log"
)

// HandlePost binds the incoming request object to a struct defined in model/message.go
func HandlePost(c *gin.Context) {
				data := &model.Message{}
				err := c.Bind(data)
				if err != nil {
								log.Fatal(err)
				}
				SaveToDB(data)
}

// SaveToDB will append a string to the messageDB variable defined in model/message.go
func SaveToDB(data *model.Message) {
				model.MessageDB["status"] = data.Status
				model.MessageDB["timestamp"] = data.Timestamp
				model.MessageDB["statusCode"] = "200"
}
