package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("src/templates/index.html"))
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", rootHandler)
	log.Println("Server starting on port 8081")
	log.Fatal("Failed to start server", http.ListenAndServe(":8081", nil))
}
