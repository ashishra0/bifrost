package handler

import (
		"fmt"
		"github.com/ashishra0/bifrost/model"
		"github.com/gin-gonic/gin"
		"log"
)

// FromHasura function handles the payload emitted by hasura event system.
func FromHasura(c *gin.Context) {
	hasura := &model.AutoGenerated{}
	err := c.Bind(hasura)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("event: %s", hasura)
}