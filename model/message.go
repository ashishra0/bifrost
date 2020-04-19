package model

type Message struct {
	Status string `json:"status"`
	Timestamp string `json:"timestamp"`
}

var MessageDB map[string] string

var ErrorDB map[string] string