package model

type Message struct {
	Status string `json:"status"`
	Timestamp string `json:"timestamp"`
}

var MessageDB = make(map[string] string)

var ErrorDB = make(map[string] string)