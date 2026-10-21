package main

import (
	"log"
	"net/http"
	"os"

	handler "example.com/learn-golang/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/health", handler.Handler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("Server running on %s", addr)
    
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
