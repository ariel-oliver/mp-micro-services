package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ariel-oliver/mp-micro-services/order-service/config"
	"github.com/ariel-oliver/mp-micro-services/order-service/models"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)

	stmt, err := config.DB.Prepare("INSERT INTO orders (user_id, product_id, quantity, total) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(order.UserID, order.ProductID, order.Quantity, order.Total)
	if err != nil {
		http.Error(w, "Could not create order", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, user_id, product_id, quantity, total FROM orders")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Total); err != nil {
			http.Error(w, "Error scanning order", http.StatusInternalServerError)
			return
		}
		orders = append(orders, order)
	}
	json.NewEncoder(w).Encode(orders)
}
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var updatedOrder models.Order
	json.NewDecoder(r.Body).Decode(&updatedOrder)

	stmt, err := config.DB.Prepare("UPDATE orders SET user_id = ?, product_id = ?, quantity = ?, total = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(updatedOrder.UserID, updatedOrder.ProductID, updatedOrder.Quantity, updatedOrder.Total, updatedOrder.ID)
	if err != nil {
		http.Error(w, "Could not update order", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedOrder)
}
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var deletedOrder models.Order
	json.NewDecoder(r.Body).Decode(&deletedOrder)

	stmt, err := config.DB.Prepare("DELETE FROM orders WHERE id = ?")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(deletedOrder.ID)
	if err != nil {
		http.Error(w, "Could not delete order", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order deleted successfully"})
}
