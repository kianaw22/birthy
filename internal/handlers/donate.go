package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/kianaw22/birthy/internal/database"
	"github.com/kianaw22/birthy/internal/zarinpal"
	"github.com/kianaw22/birthy/models"
)

// DonateHandler handles donations and redirects to Zarinpal
func DonateHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Parse donation amount
	amountStr := r.FormValue("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil || amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	// Request payment from Zarinpal
	authority, err := zarinpal.RequestPayment(amount)
	if err != nil {
		http.Error(w, "Error contacting Zarinpal: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Save donation in the database
	donation := models.Donation{
		Amount:    amount,
		Status:    "pending",
		Authority: authority,
	}
	if err := database.GetDB().WithContext(ctx).Create(&donation).Error; err != nil {
		http.Error(w, "Error saving donation", http.StatusInternalServerError)
		return
	}

	// Redirect to Zarinpal payment page
	http.Redirect(w, r, "https://www.zarinpal.com/pg/StartPay/"+authority, http.StatusSeeOther)
}
