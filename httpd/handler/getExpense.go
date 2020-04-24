package handler

import (
		"fmt"
		"github.com/ashishra0/raspberry-pi-service/query"
		"github.com/gin-gonic/gin"
)

// GetExpense calls the FetchFromHasura function and sends the
// result returned from that function to the client.
func GetExpense(c *gin.Context) {
		record := query.FetchFromHasura()
		// @todo: Send record obtained from hasura to the client
		for _, name := range record.Expenses {
				fmt.Println(name)
		}
}
