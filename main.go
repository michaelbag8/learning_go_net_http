package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Home Page"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("About Page"))
}

func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Contact Page"))
}

func greetHandler(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")

	lang := r.URL.Query().Get("lang")

	if name == ""{
		name = "Guest"
	}

	if lang == "es"{
		fmt.Fprintf(w, "Hola %s", name)
	}else{
		fmt.Fprintf(w, "Hello %s", name)
	}
	
}

func farewellHandler(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")

	lang := r.URL.Query().Get("lang")

	if name == ""{
		name = "Guest"
	}

	if lang == "es"{
		fmt.Fprintf(w, "Adiós, %s", name)
	}else{
		fmt.Fprintf(w, "Goodbye, %s", name)
	}
}


func getUserIdHandler(w http.ResponseWriter, r *http.Request){
	path :=r.URL.Path

	parts := strings.Split(path, "/")

	if len(parts) < 3{
		http.Error(w, "User ID not provided", http.StatusBadRequest)
		return
	}

	userID := parts[2]
	fmt.Fprintf(w, "User ID: %s", userID)
}



func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/farewell", farewellHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/users/", getUserIdHandler)

	if err := http.ListenAndServe(":9090", nil); err != nil{
		log.Fatal("Server failed to start: ",err)
	}
}