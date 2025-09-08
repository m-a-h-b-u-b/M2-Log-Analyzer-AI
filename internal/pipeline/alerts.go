//! M2 Log Analyzer AI
//! --------------------------------
//! License : Dual License
//!           - Apache 2.0 for open-source / personal use
//!           - Commercial license required for closed-source use
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:


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
