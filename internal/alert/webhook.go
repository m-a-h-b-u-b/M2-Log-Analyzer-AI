//! Module Name: webhook.go
//! --------------------------------
//! License : Apache 2.0
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Generic webhook integration for sending anomaly alerts.

package alert

import (
	"bytes"
	"net/http"
)

type WebhookAlert struct {
	URL string
}

func NewWebhookAlert(url string) *WebhookAlert {
	return &WebhookAlert{URL: url}
}

func (w *WebhookAlert) Send(message string) error {
	_, err := http.Post(w.URL, "text/plain", bytes.NewBuffer([]byte(message)))
	return err
}
