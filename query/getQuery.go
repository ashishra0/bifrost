package query

import (
	"log"

	"github.com/ashishra0/bifrost/model"
	"github.com/shahidhk/gql"
)

// GetQuery sends a query to hasura to fetch
// all values from the db.
func GetQuery() model.Expense {
	client := gql.NewClient("https://alfred-expense-beta.herokuapp.com/v1/graphql", nil)
	var record model.Expense
	err := client.Execute(gql.Request{Query: `query {
  Expense (order_by: {item_id: desc}) {
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
