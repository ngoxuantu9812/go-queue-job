package worker

import (
	"log"

	"demo1/config"
	"demo1/tasks"
	"github.com/hibiken/asynq"
)

// StartWorkerServer khởi chạy worker server để xử lý các job trong hàng đợi
func StartWorkerServer() {
	// Tạo server với cấu hình Redis từ config
	server := asynq.NewServer(config.RedisConfig(), asynq.Config{
		Concurrency: 10, // Số lượng worker chạy song song
	})

	// Đăng ký các handler cho các loại job
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.EmailTask, tasks.HandleEmailTask)

	// Chạy server để xử lý các job
	if err := server.Run(mux); err != nil {
		log.Fatalf("Không thể chạy worker: %v", err)
	}
}
