// jobs.go
package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

// Payload for the email delivery tasks
type EmailDeliveryPayload struct {
	To      string
	Subject string
	Body    string
}

// CreateEmailDeliveryTask creates a new email delivery tasks with the given payload.
func CreateEmailDeliveryTask(to, subject, body string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{To: to, Subject: subject, Body: body})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(EmailTask, payload), nil
}
