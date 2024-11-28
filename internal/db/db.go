package db

import (
	"database/sql"
)

// CreateTable ensures the URLs exists
func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        short_url TEXT NOT NULL,
        original_url TEXT NOT NULL        
    );`
	_, err := db.Exec(query)
	return err
}

// StoreURL stores a new short and original URL in the database
func StoreURL(db *sql.DB, shortURL, originalURL string) error {
	query := `INSERT INTO urls (short_url, original_url) VALUES (?, ?)`
	_, err := db.Exec(query, shortURL, originalURL)
	return err
}

// GetOriginalURL retrieves the original URL by its short URL
func GetOriginalURL(db *sql.DB, shortURL string) (string, error) {
	query := `SELECT original_url FROM urls WHERE short_url = ?`
	row := db.QueryRow(query, shortURL)
	var originalURL string
	err := row.Scan(&originalURL)
	return originalURL, err
}