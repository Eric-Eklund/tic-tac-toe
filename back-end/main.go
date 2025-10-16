package main

import (
	"fmt"
	"log"
	"net/http"
	_ "path/filepath"
)

func main() {
	fmt.Println("Tic-tac-toe backend starting...")

	// API routes
	http.HandleFunc("/api/health", healthHandler)
	// Add more API routes here later

	// Serve static files (React build)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status":"ok","service":"tic-tac-toe-backend"}`)
}
