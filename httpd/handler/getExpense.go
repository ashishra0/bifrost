package handler

import (
		"github.com/ashishra0/raspberry-pi-service/query"
		"github.com/gin-gonic/gin"
		"net/http"
)

// GetExpense calls the GetQuery function and sends the
// result returned from that function to the client.
func GetExpense(c *gin.Context) {
	record := query.GetQuery()
	// @todo: Send record obtained from hasura to the client
	c.JSON(http.StatusOK, gin.H{"data": record})
}
