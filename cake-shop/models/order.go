package models

import "time"

type OrderItem struct {
	CakeID int     `json:"cakeId"`
	Name   string  `json:"name"`
	Size   string  `json:"size"`
	Price  float64 `json:"price"`
}

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Order struct {
	ID           string      `json:"id"`
	Customer     Customer    `json:"customer"`
	Items        []OrderItem `json:"items"`
	Subtotal     float64     `json:"subtotal"`
	Tax          float64     `json:"tax"`
	Total        float64     `json:"total"`
	DeliveryDate string      `json:"deliveryDate"`
	Instructions string      `json:"instructions"`
	CreatedAt    time.Time   `json:"createdAt"`
	Status       string      `json:"status"`
}
