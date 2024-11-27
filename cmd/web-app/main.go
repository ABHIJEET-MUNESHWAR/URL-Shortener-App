package main

import (
	"github.com/ABHIJEET-MUNESHWAR/URL-Shortener-App/internal/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.ShowIndex)
	http.HandleFunc("/shorten", controllers.Shorten)
	log.Fatal(http.ListenAndServe(":8080", nil))
}