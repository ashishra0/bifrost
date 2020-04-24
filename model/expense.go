package model

import "time"

// AutoGenerated is a collection of response sent by the
// hasura event system.
type AutoGenerated struct {
	Event        Event        `json:"event"`
	CreatedAt    time.Time    `json:"created_at"`
	ID           string       `json:"id"`
	DeliveryInfo DeliveryInfo `json:"delivery_info"`
	Trigger      Trigger      `json:"trigger"`
	Table        Table        `json:"table"`
}

// SessionVariables contains user role.
type SessionVariables struct {
	XHasuraRole string `json:"x-hasura-role"`
}

// New contains the table field names present in the db.
type New struct {
	ItemID   int    `json:"item_id"`
	ItemName string `json:"item_name"`
	ItemCost int    `json:"item_cost"`
}

// Data is a collection of Old and New struct.
type Data struct {
	Old interface{} `json:"old"`
	New New         `json:"new"`
}

// Event is a collection of session variables and data struct.
type Event struct {
	SessionVariables SessionVariables `json:"session_variables"`
	Op               string           `json:"op"`
	Data             Data             `json:"data"`
}

// DeliveryInfo provides information about the payload delivery.
type DeliveryInfo struct {
	MaxRetries   int `json:"max_retries"`
	CurrentRetry int `json:"current_retry"`
}

// Trigger provides the name of the event trigger.
type Trigger struct {
	Name string `json:"name"`
}

// Table provides the details about the database in hasura.
type Table struct {
	Schema string `json:"schema"`
	Name   string `json:"name"`
}

// Expense is a struct having a field of type Expenses
// Field Expenses is a slice of type New
// It gets appended a new expenses whenever there is a
// POST request with data sent to /api/expense/create
type Expense struct {
	Expenses []New `json:"Expense"`
}
