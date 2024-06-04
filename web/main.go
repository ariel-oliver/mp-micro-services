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
	router.HandleFunc("/profile", profileHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
func profileHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "profile.html", nil)
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

		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "profile.html", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {

	req, _ := http.NewRequest("GET", "http://localhost:8081/products", nil)

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
