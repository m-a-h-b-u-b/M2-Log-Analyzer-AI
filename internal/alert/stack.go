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
//! Slack integration for sending anomaly alerts.

package alert

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SlackAlert struct {
	WebhookURL string
}

func NewSlackAlert(url string) *SlackAlert {
	return &SlackAlert{WebhookURL: url}
}

func (s *SlackAlert) Send(message string) error {
	payload := map[string]string{"text": message}
	data, _ := json.Marshal(payload)

	_, err := http.Post(s.WebhookURL, "application/json", bytes.NewBuffer(data))
	return err
}
