package pipeline

import (
	"bytes"
	"log"
	"net/http"
)

func alertSlack(webhook, msg string) {
	payload := []byte(`{"text":"` + msg + `"}`)
	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Slack alert error: %v", err)
		return
	}
	resp.Body.Close()
}

func alertWebhook(webhook, msg string) {
	payload := []byte(msg)
	resp, err := http.Post(webhook, "text/plain", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Webhook alert error: %v", err)
		return
	}
	resp.Body.Close()
}
