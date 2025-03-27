package main

import (
	"log"
	"net/http"

	"github.com/kianaw22/birthy/internal/database"
	"github.com/kianaw22/birthy/internal/handlers"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Serve static files (CSS, JS, images, etc.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/donate", handlers.DonateHandler)
	http.HandleFunc("/verify", handlers.VerifyHandler)

	// Start the server
	log.Println("✅ Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
