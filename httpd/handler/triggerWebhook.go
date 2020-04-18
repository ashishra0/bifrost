package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func TriggerWebhook(c * gin.Context) {
	useless := map[string]string {
		"status": "ok",
	}

	_, err := http.Post("https://smee.io/mQAzpDaQKOQpXyr0", "Application/json", bytes.NewBuffer(jsonMarshal(useless)))
	if err != nil {
		log.Fatalln(err)
	}
}

func jsonMarshal(v interface{}) []byte {
	val, err := json.Marshal(v)
	if err != nil {
		log.Fatalln(err)
	}
	return val
}