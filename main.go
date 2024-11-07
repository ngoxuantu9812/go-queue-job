// worker.go
package main

import (
	"context"
	"demo1/tasks"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	go createServerAsync()
	createServerHttp()

}

func createServerHttp() {
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring", // RootPath specifies the root for asynqmon app
		RedisConnOpt: asynq.RedisClientOpt{Addr: ":6379"},
	})
	r := mux.NewRouter()
	r.HandleFunc("/add-job-to-queue", AddJobToQueue).Methods("GET")
	r.PathPrefix(h.RootPath()).Handler(h)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
		BaseContext: func(listener net.Listener) context.Context {
			ctx = context.WithValue(ctx, "serverAddr", listener.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}

}

func createServerAsync() {
	// Connect to Redis server
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 10, // Number of workers
		},
	)
	// Define a handler for the email delivery job type
	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.EmailTask, HandleEmailDeliveryTask)
	// Start processing jobs
	if err := server.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
func AddJobToQueue(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	go tasks.EnqueueJob()

}

// HandleEmailDeliveryTask processes the email delivery tasks
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	startTime := time.Now()

	var p tasks.EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("could not unmarshal payload: %v", err)
	}
	// Simulate sending an email
	fmt.Printf("Sending email to %s with subject %q and body %q\n", p.To, p.Subject, p.Body)
	elapsed := time.Since(startTime)
	log.Printf("Task %s hoàn thành trong: %v", t, elapsed)

	return nil
}
