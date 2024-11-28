package main

import (
	"database/sql"
	"github.com/ABHIJEET-MUNESHWAR/URL-Shortener-App/internal/controllers"
	"github.com/ABHIJEET-MUNESHWAR/URL-Shortener-App/internal/db"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	sqlite, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite.Close()

	if err := db.CreateTable(sqlite); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {
			controllers.ShowIndex(writer, request)
		} else {
			controllers.Proxy(sqlite)(writer, request)
		}
	})

	http.HandleFunc("/shorten", controllers.Shorten(sqlite))
	log.Fatal(http.ListenAndServe(":8080", nil))
}