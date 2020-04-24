package handler

import (
				"encoding/json"
				"github.com/ashishra0/raspberry-pi-service/model"
				"github.com/ashishra0/raspberry-pi-service/query"
				"github.com/gin-gonic/gin"
				"io/ioutil"
				"log"
)

// CreateExpense is handler function that will take the request body
// and unmarshal it to model.New struct defined in model/expense
// then It will call a mutation function and pass the unmarshaled struct
// as argument.
func CreateExpense(c *gin.Context) {
				var record model.New
				b, err := ioutil.ReadAll(c.Request.Body)
				if err != nil {
								log.Fatal(err)
				}
				err = json.Unmarshal(b, &record)
				if err != nil {
								log.Fatal(err)
				}
				SendMutation(record)
}

// SaveVar will add new record in the []New and
// return a map[string]interface{} of that record.
func SaveVar(expense model.New) map[string]interface{} {
				return map[string]interface{}{
								"item_name": expense.ItemName,
								"item_cost": expense.ItemCost,
				}
}

// SendMutation saves the return type of SaveVar function
// and save it as a variable and pass it to SendToHasura() func
// which is responsible for sending the actual mutation to the graphql
// server.
func SendMutation(v model.New) {
				variables := SaveVar(v)
				query.SendToHasura(variables)
}
