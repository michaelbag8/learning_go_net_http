package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ─── Helper ──────────────────────────────────────────
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// ─── Structs ─────────────────────────────────────────
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Contact struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// ─── Handlers ────────────────────────────────────────
func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed", // ✅
		})
		return
	}

	query := r.URL.Query().Get("q")
	limit := r.URL.Query().Get("limit")

	writeJSON(w, http.StatusOK, map[string]string{
		"query":   query,
		"limit":   limit,
		"message": "showing results for " + query,
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed", // ✅
		})
		return
	}

	defer r.Body.Close()

	var req Login
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	// ✅ Use req fields directly — no unnecessary variables
	if req.Email == "" || req.Password == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "email and password are required",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "login successful",
		"email":   req.Email,
	})
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "method not allowed", // ✅
		})
		return
	}

	defer r.Body.Close()

	var req Contact
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	// ✅ Use req fields directly — no unnecessary variables
	if req.Name == "" || req.Subject == "" || req.Message == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "name, subject and message are required",
		})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{
		"message": "message sent successfully",
		"name":    req.Name,
	})
}

// ─── Main ────────────────────────────────────────────
func main() {
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Server is running at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}