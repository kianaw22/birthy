package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kianaw22/birthy/internal/database"
	"github.com/kianaw22/birthy/internal/zarinpal"
	"github.com/kianaw22/birthy/models"
)

// VerifyHandler verifies the payment status from Zarinpal
func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	authority := r.URL.Query().Get("Authority")
	status := r.URL.Query().Get("Status")

	if authority == "" || status == "" {
		http.Error(w, "Invalid request parameters", http.StatusBadRequest)
		return
	}

	// Check if payment was successful
	if status != "OK" {
		updateDonationStatus(ctx, authority, "failed", 0)
		fmt.Fprint(w, "Payment was cancelled")
		return
	}

	var donation models.Donation
	result := database.GetDB().WithContext(ctx).Where("authority = ?", authority).First(&donation)
	if result.Error != nil || donation.Status != "pending" {
		http.Error(w, "Donation not found or already processed", http.StatusBadRequest)
		return
	}

	// Verify payment with Zarinpal
	ok, refID, err := zarinpal.VerifyPayment(authority, donation.Amount)
	if err != nil || !ok {
		updateDonationStatus(ctx, authority, "failed", 0)
		http.Error(w, "Payment not verified", http.StatusInternalServerError)
		return
	}

	// Update donation status
	updateDonationStatus(ctx, authority, "success", refID)
	fmt.Fprintf(w, "Payment successful! Tracking code: %d", refID)
}

// updateDonationStatus updates the status and ref_id of the donation
func updateDonationStatus(ctx context.Context, authority string, status string, refID int64) {
	database.GetDB().WithContext(ctx).
		Model(&models.Donation{}).
		Where("authority = ?", authority).
		Updates(map[string]interface{}{
			"status": status,
			"ref_id": refID,
		})
}
