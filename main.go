package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Contact Page"))
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	lang := r.URL.Query().Get("lang")

	if name == "" {
		name = "Guest"
	}

	if lang == "es" {
		fmt.Fprintf(w, "Hola %s", name)
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}

}

func farewellHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	lang := r.URL.Query().Get("lang")

	if name == "" {
		name = "Guest"
	}

	if lang == "es" {
		fmt.Fprintf(w, "Adiós, %s", name)
	} else {
		fmt.Fprintf(w, "Goodbye, %s", name)
	}
}

func getUserIdHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Expected format: /users/{id}", http.StatusBadRequest)
		return
	}

	userID := parts[2]

	switch userID {
	case "123":
		fmt.Fprintf(w, "User ID: %s, Name: Michael", userID)
	case "456":
		fmt.Fprintf(w, "User ID: %s, Name: Sarah", userID)
	default:
		fmt.Fprintf(w, "User ID: %s, Name: Unknown", userID)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var users []string

	if r.Method != "POST" {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	type User struct {
		Name string `json:"name"`
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	type Response struct {
		Message string `json:"message"`
	}

	users = append(users, user.Name)

	resp := Response{Message: fmt.Sprintf("User %s created", user.Name)}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	method := r.Method

	switch method {
	case "POST":
		fmt.Fprintln(w, `{"message": "Creating a user"}`)
	case "GET":
		fmt.Fprintln(w, `{"message": "Listing all users"}`)

	case "PUT":
		fmt.Fprintln(w, `{"message": "Updating a user"}`)

	case "DELETE":
		fmt.Fprintln(w, `{"message": "Deleting a user"}`)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/farewell", farewellHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/users/", getUserIdHandler)
	http.HandleFunc("/user", getUsers)
	http.HandleFunc("/users/create", createUser)

	fmt.Println("server is runing ......http://localhost:9090")

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
