package tasks

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

// EmailTaskPayload là payload cho job gửi email
type EmailTaskPayload struct {
	To      string
	Subject string
	Body    string
}

// NewEmailTask tạo job gửi email với payload được mã hóa
func NewEmailTask(to, subject, body string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailTaskPayload{To: to, Subject: subject, Body: body})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(EmailTask, payload), nil
}
