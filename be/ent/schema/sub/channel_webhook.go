package sub

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/echoturing/log"
	"github.com/labstack/echo"
)

type Webhook struct {
	URL string `json:"url"`
}

type WebhookMessage struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func (w *Webhook) SendWebhookAlert(ctx context.Context, msg *WebhookMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	// TODO: temp ignore response
	log.DebugWithContext(ctx, "send post", "data", string(data), "url", w.URL)
	_, err = http.DefaultClient.Post(w.URL, echo.MIMEApplicationJSON, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	return nil
}
