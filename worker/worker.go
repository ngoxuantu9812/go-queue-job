package worker

import (
	"log"

	"demo1/config"
	"demo1/tasks"
	"github.com/hibiken/asynq"
)

func StartWorkerServer() {
	server := asynq.NewServer(config.RedisConfig(), asynq.Config{
		Concurrency: 8,
	})
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.EmailTask, tasks.HandleEmailTask)
	// Chạy server để xử lý các job
	if err := server.Run(mux); err != nil {
		log.Fatalf("Không thể chạy worker: %v", err)
	}

	log.Println("Worker start")
}
