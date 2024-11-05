package main

import (
	"log"

	"demo1/config"
	"demo1/tasks"
	"demo1/worker"
	"github.com/hibiken/asynq"
)

func main() {
	// Khởi động worker server ở một goroutine khác
	go worker.StartWorkerServer()

	// Kết nối tới Redis server và tạo client
	client := asynq.NewClient(config.RedisConfig())
	defer client.Close()

	// Tạo job gửi email
	task, err := tasks.NewEmailTask("user@example.com", "Welcome", "Thanks for signing up!")
	if err != nil {
		log.Fatalf("Không thể tạo job gửi email: %v", err)
	}

	// Thêm job vào hàng đợi với các tùy chọn
	info, err := client.Enqueue(task, asynq.Queue("critical"))
	if err != nil {
		log.Fatalf("Không thể thêm job vào hàng đợi: %v", err)
	}

	log.Printf("Đã thêm job vào hàng đợi với ID: %s", info.ID)

	// Giữ main goroutine hoạt động để worker chạy liên tục
	select {}
}
