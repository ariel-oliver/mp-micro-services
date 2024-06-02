package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	store = sessions.NewCookieStore([]byte("something-very-secret"))
	tpl   = template.Must(template.ParseGlob("templates/*"))
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/register", registerHandler).Methods("GET", "POST")
	router.HandleFunc("/login", loginHandler).Methods("GET", "POST")
	router.HandleFunc("/products", productsHandler).Methods("GET")
	router.HandleFunc("/orders", ordersHandler).Methods("GET", "POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user := map[string]string{
			"username": username,
			"password": password,
		}
		jsonData, _ := json.Marshal(user)

		resp, err := http.Post("http://localhost:8083/register", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Error registering user", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Failed to register", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "register.html", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		user := map[string]string{
			"username": username,
			"password": password,
		}
		jsonData, _ := json.Marshal(user)

		resp, err := http.Post("http://localhost:8083/login", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, "Error logging in", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		var result map[string]string
		json.NewDecoder(resp.Body).Decode(&result)
		token := result["token"]

		session, _ := store.Get(r, "session")
		session.Values["token"] = token
		session.Save(r, w)

		http.Redirect(w, r, "/products", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	token, ok := session.Values["token"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	req, _ := http.NewRequest("GET", "http://localhost:8081/products", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	var products []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&products)

	tpl.ExecuteTemplate(w, "products.html", products)
}

func ordersHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	token, ok := session.Values["token"].(string)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		userID := r.FormValue("user_id")
		productID := r.FormValue("product_id")
		quantity := r.FormValue("quantity")
		total := r.FormValue("total")

		order := map[string]string{
			"user_id":    userID,
			"product_id": productID,
			"quantity":   quantity,
			"total":      total,
		}
		jsonData, _ := json.Marshal(order)

		req, _ := http.NewRequest("POST", "http://localhost:8082/orders", bytes.NewBuffer(jsonData))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "Error creating order", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Failed to create order", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/orders", http.StatusSeeOther)
		return
	}

	req, _ := http.NewRequest("GET", "http://localhost:8082/orders", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error fetching orders", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	var orders []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&orders)

	tpl.ExecuteTemplate(w, "orders.html", orders)
}
