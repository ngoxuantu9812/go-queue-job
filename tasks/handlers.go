package tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

// HandleEmailTask xử lý job gửi email
func HandleEmailTask(ctx context.Context, t *asynq.Task) error {
	var p EmailTaskPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("không thể parse payload: %w", err)
	}

	// Thực hiện gửi email (giả lập)
	fmt.Printf("Gửi email đến %s với subject: %s và body: %s\n", p.To, p.Subject, p.Body)
	return nil
}
