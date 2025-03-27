package main

import (
	"log"
	"net/http"

	"github.com/kianaw22/birthy/internal/database"
	"github.com/kianaw22/birthy/internal/handlers"
	"github.com/kianaw22/birthy/models"
)

func runMigrations() {
	db := database.GetDB()
	err := db.AutoMigrate(&models.Donation{})
	if err != nil {
		log.Fatalf("❗️ Error running migrations: %v", err)
	}
	log.Println("✅ Migrations applied successfully!")
}

func main() {
	// Initialize the database
	database.InitDB()
	runMigrations()
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
