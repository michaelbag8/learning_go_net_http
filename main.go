package main

import (
	"fmt"
	"log"
	"net/http"
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
	name := r.URL.Query().Get("param")

	lang := r.URL.Query().Get("param")

	if name == ""{
		name = "Guest!"
	}

	if lang == "es"{
		fmt.Fprintf(w, "Halo %v", name)
	}else{
		fmt.Fprintf(w, "Hello %v", name)
	}
	
}

func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	http.HandleFunc("/greet", greetHandler)

	if err := http.ListenAndServe(":9090", nil); err != nil{
		log.Fatal("Server failed to start: ",err)
	}
}