package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", ShowHomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func ShowHomePage(w http.ResponseWriter, r *http.Request) {
	template, error := template.ParseFiles("internal/views/index.html")
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
	if error = template.Execute(w, nil); error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}
}