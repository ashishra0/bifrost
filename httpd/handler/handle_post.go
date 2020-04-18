package handler

import (
	"github.com/ashishra0/raspberry-pi-service/model"
	"github.com/gin-gonic/gin"
	"log"
)

func HandlePost(c *gin.Context) {
	data := &model.Message{}
	err := c.Bind(data)
	if err != nil {
		log.Fatal(err)
	}
	saveToDB(data.Status)
}

//@todo - save to DB
func saveToDB(data string) {
	model.MessageDB = data
}