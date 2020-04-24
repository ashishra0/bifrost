package query

import (
	"github.com/ashishra0/raspberry-pi-service/model"
	"github.com/shahidhk/gql"
	"log"
)

// SendToHasura sends a insert mutation to a graphql client
func SendToHasura(v map[string]interface{}) {
	client := gql.NewClient("http://localhost:8080/v1/graphql", nil)
	var expenses = model.New{}
	err := client.Execute(gql.Request{Query: `mutation ($item_name: String, $item_cost: Int){
		insert_Expense( objects: [ { item_name: $item_name, item_cost: $item_cost } ] ) { returning { item_id }}
	}`, Variables: v}, &expenses)
	if err != nil {
		log.Fatal(err)
	}
}
