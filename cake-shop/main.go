package main

import (
	"cake-shop/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static frontend files (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// API endpoints
	http.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS for local development
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		switch r.Method {
		case http.MethodPost:
			handlers.CreateOrder(w, r)
		case http.MethodGet:
			handlers.GetOrders(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	port := "8080"
	fmt.Printf("🍰 Cake Shop running at http://localhost:%s\n", port)
	fmt.Printf("📋 View orders at http://localhost:%s/api/orders\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
