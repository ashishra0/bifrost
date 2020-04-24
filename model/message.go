package model

// Message defines the structure of the payload.
type Message struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// MessageDB is a key value pair of the message payload.
var MessageDB = make(map[string]string)

// ErrorDB is a key value pair of the error payload.
var ErrorDB = make(map[string]string)
