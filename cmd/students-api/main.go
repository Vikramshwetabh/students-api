package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"githubh.com/vikramshwetabh/students-api/internal/http/handlers/student"
)

func main() {
	// Load configuration
	cfg := struct {
		Addr string
	}{
		Addr: ":8080",
	}

	// Setup router
	router := http.NewServeMux()
	router.HandleFunc("POST/api/students", student.New()) // Register the student handler

	// Setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	// Handle graceful shutdown
	slog.Info("Starting server...", slog.String("address", cfg.Addr))
	// Create a channel to listen for interrupt signals
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // Catch interrupt signal

	// Start server in a goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server: ", err)
		}
	}()
	<-done // Wait for a signal to stop the server

	slog.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Create a context with timeout for graceful shutdown
	defer cancel()

	err := server.Shutdown(ctx) // Shutdown the server gracefully
	if err != nil {
		slog.Error("Failed to shutdown: ", slog.String("error", err.Error()))
	} else {
		slog.Info("Server stopped Successfully")
	}
}
