// enqueue_job.go
package tasks

import (
	"demo1/config"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

func EnqueueJob() {
	// Connect to Redis server
	client := asynq.NewClient(config.RedisConfig())
	defer client.Close()
	fmt.Println("EnqueueJob")
	// Create a new email delivery tasks
	task, err := CreateEmailDeliveryTask("user@example.com", "Welcome!", "Hello and welcome!")
	if err != nil {
		log.Fatalf("could not create tasks: %v", err)
	}
	// Enqueue the tasks
	for i := 0; i < 1000; i++ {
		_, _ = client.Enqueue(task)

	}

	if err != nil {
		log.Fatalf("could not enqueue tasks: %v", err)
	}

	//log.Printf("Enqueued tasks: id=%s queue=%s", info.ID, info.Queue)
}
