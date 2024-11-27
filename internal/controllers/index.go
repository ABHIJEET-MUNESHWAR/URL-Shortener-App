package controllers

import (
	"html/template"
	"net/http"
)

func ShowIndex(w http.ResponseWriter, r *http.Request) {
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