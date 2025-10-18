package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"todo-api/internal/todo"
)

func main() {
	store := todo.NewInMemoryStore()
	h := todo.NewHandler(store)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	handler := todo.LoggingMiddleware(mux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// graceful shutdown
	idleConnsClosed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		fmt.Println("\nshutting down server...")
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	fmt.Println("listening on http://localhost:8080")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	fmt.Println("server stopped")
}
