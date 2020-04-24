package helper

import (
	"github.com/ashishra0/raspberry-pi-service/model"
	"time"
)

// GetTime is a helper function that returns the current time
// In "DD-MM-YYYY H:M:S DAY" format.
func GetTime() string {
	dt := time.Now()

	return dt.Format("01-02-2006 15:04:05 Monday")
}

// SetError is a helper function that assigns values to model.ErrorDB map.
func SetError() {
	model.ErrorDB["status"] = "No power"
	model.ErrorDB["timestamp"] = GetTime()
	model.ErrorDB["statusCode"] = "404"
}

// GetError is a helper function that returns model.ErrorDB map.
func GetError() map[string]string {
	SetError()
	return model.ErrorDB
}

// SetDefaultStatus is a helper function that sets default values to model.MessageDB map after each GET request.
func SetDefaultStatus() {
	model.MessageDB["status"] = "No response from Tesla"
	model.MessageDB["timestamp"] = GetTime()
	model.MessageDB["statusCode"] = ""
}
