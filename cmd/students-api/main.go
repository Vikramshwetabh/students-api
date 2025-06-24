package main

import (
	"fmt"
	"log"
	"net/http"
	//	"github.com/vikramshwetabh/students-api/internal/config"
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
	router.HandleFunc("Get/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Students API!"))
	})

	// Setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("Server is running on", cfg.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
