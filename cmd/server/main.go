package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/developer-experience-DevEX-platform/authentication-api/internal/app"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           app.New("authentication-api").Handler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Server listening on http://0.0.0.0:%s", port)
	log.Fatal(server.ListenAndServe())
}
