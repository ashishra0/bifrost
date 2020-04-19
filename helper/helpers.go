package helper

import (
	"github.com/ashishra0/raspberry-pi-service/model"
	"time"
)


func GetTime() string {
	dt := time.Now()

	return dt.Format("01-02-2006 15:04:05 Monday")
}

func SetError() {
	model.ErrorDB["status"] = "No power"
	model.ErrorDB["timestamp"] = GetTime()
}

func GetError() map[string]string {
	SetError()
	return model.ErrorDB
}

func SetDefaultStatus() {
	model.MessageDB["status"] = "No response from Tesla"
	model.MessageDB["TimeStamp"] = GetTime()
}