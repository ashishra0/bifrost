package query

import (
	"github.com/ashishra0/bifrost/model"
	"github.com/shahidhk/gql"
	"log"
)

// InsertQuery sends a insert mutation to a graphql client
func InsertQuery(v map[string]interface{}) {
	client := gql.NewClient("https://alfred-expense-beta.herokuapp.com/v1/graphql", nil)
	var expenses = model.New{}
	err := client.Execute(gql.Request{Query: `mutation ($item_name: String, $item_cost: Int){
		insert_Expense( objects: [ { item_name: $item_name, item_cost: $item_cost } ] ) { returning { item_id }}
	}`, Variables: v}, &expenses)
	if err != nil {
		log.Fatal(err)
	}
}
