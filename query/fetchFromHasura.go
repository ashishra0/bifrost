package query

import (
		"github.com/ashishra0/raspberry-pi-service/model"
		"github.com/shahidhk/gql"
		"log"
)

// FetchFromHasua sends a query to hasura to fetch
// all values from the db.
func FetchFromHasura() model.Expense {
		client := gql.NewClient("http://localhost:8080/v1/graphql", nil)
		var record model.Expense
		err := client.Execute(gql.Request{Query: `query {
  Expense {
    item_name
  }
}`}, &record)
		if err != nil {
				log.Fatal(err)
		}
		return record
}
