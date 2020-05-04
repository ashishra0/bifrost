package query

import (
	"github.com/ashishra0/raspberry-pi-service/model"
	"github.com/shahidhk/gql"
	"log"
)

// GetQuery sends a query to hasura to fetch
// all values from the db.
func GetQuery() model.Expense {
	client := gql.NewClient("http://localhost:8080/v1/graphql", nil)
	var record model.Expense
	err := client.Execute(gql.Request{Query: `query {
  Expense (order_by: {date: desc}) {
    item_id
    item_name,
    item_cost,
		date
  }
}`}, &record)
	if err != nil {
		log.Fatal(err)
	}
	return record
}
