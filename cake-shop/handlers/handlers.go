package handlers

import (
	"cake-shop/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

const ordersFile = "orders.json"

var (
	mu     sync.Mutex
	orders []models.Order
)

// Load existing orders from file on startup
func init() {
	data, err := os.ReadFile(ordersFile)
	if err == nil {
		json.Unmarshal(data, &orders)
	}
}

// Save orders to JSON file
func saveOrders() error {
	data, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(ordersFile, data, 0644)
}

// POST /api/orders - Create a new order
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate order ID and timestamp
	order.ID = fmt.Sprintf("ORD-%d", time.Now().UnixNano())
	order.CreatedAt = time.Now()
	order.Status = "Pending"

	mu.Lock()
	orders = append(orders, order)
	if err := saveOrders(); err != nil {
		mu.Unlock()
		http.Error(w, "Failed to save order", http.StatusInternalServerError)
		return
	}
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Order placed successfully!",
		"order":   order,
	})
}

// GET /api/orders - View all orders (for admin)
func GetOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
