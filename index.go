package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ─── Structs ─────────────────────────────────────────
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// ─── Helper ──────────────────────────────────────────
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ─── Handlers ────────────────────────────────────────
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:    78,
		Name:  "Michael Bag",
		Email: "michaelbag8@gmail.com",
		Role:  "super admin",
	}
	writeJSON(w, http.StatusOK, user)
}

func missingHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusNotFound, map[string]string{
		"error": "resource not found",
	})
}

func createdHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusCreated, map[string]any{
		"message": "user created successfully", 
		"id":      345,
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	language := r.URL.Query().Get("language")

	if name == "" {
		name = "World"
	}

	switch strings.ToLower(language) { 
	case "french":
		fmt.Fprintf(w, "Bonjour, %s!\n", name)
	case "spanish":
		fmt.Fprintf(w, "Hola, %s!\n", name)
	default:
		fmt.Fprintf(w, "Hello, %s!\n", name)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page")
}

// ─── Main ────────────────────────────────────────────
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/user", getUserHandler)
	http.HandleFunc("/missing", missingHandler)
	http.HandleFunc("/created", createdHandler)

	fmt.Println("Server is running at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}