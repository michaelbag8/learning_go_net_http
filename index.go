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

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

func profileHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		username = "Guest"
	}
	appVersion := r.Header.Get("X-App-Version")

	
		writeJSON(w, http.StatusOK, map[string]string{
			"username":    username,
			"app_version": appVersion,
			"message":     "profile fetched successfully",
		})
		
	}


func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed",
		})
		return

	}

	defer r.Body.Close()

	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "all fields are required",
		})
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{
		"message": "registration successful",
		"name":    req.Name,
		"email":   req.Email,
	})

}

// ─── Main ────────────────────────────────────────────
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/user", getUserHandler)
	http.HandleFunc("/missing", missingHandler)
	http.HandleFunc("/created", createdHandler)
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("Server is running at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
