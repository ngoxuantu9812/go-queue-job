package main

import (
	"demo1/tasks"
	"log"
	"net/http"
	"time"

	"demo1/config"
	"demo1/worker"
	"github.com/hibiken/asynq"
)

var client = asynq.NewClient(config.RedisConfig())

func main() {
	go worker.StartWorkerServer()
	defer client.Close()
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(addTask))
	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}

}

func addTask(w http.ResponseWriter, r *http.Request) {

	task, err := tasks.NewEmailTask("user@example.com", "Welcome", "Thanks for signing up!")
	if err != nil {
		log.Fatalf("Không thể tạo payload cho job: %v", err)
	}

	// Tạo một task và thêm vào hàng đợi
	info, err := client.Enqueue(task, asynq.Queue("critical"), asynq.MaxRetry(3), asynq.Timeout(30*time.Second))
	if err != nil {
		log.Fatalf("Không thể thêm job vào hàng đợi: %v", err)
	}

	log.Printf("Đã thêm job vào hàng đợi với ID: %s", info.ID)
}
