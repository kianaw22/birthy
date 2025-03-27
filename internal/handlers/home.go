package handlers

import (
	"context"
	"html/template"
	"net/http"
	"time"

	"github.com/kianaw22/birthy/config"
	"github.com/kianaw22/birthy/internal/database"
	"github.com/kianaw22/birthy/models"
)

// HomeHandler renders the homepage with donation progress
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Calculate total donations using GORM
	var total int
	result := database.GetDB().WithContext(ctx).
		Model(&models.Donation{}).
		Select("COALESCE(SUM(amount), 0)").
		Where("status = ?", "success").
		Scan(&total)

	if result.Error != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	// Get target amount from config
	target := config.AppConfig.TargetAmount
	progress := float64(total) / float64(target) * 100

	data := map[string]interface{}{
		"TotalAmount":  total,
		"TargetAmount": target,
		"Progress":     progress,
	}

	tmpl.Execute(w, data)
}
