package zarinpal

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/kianaw22/birthy/config"
)

// ZarinpalRequest represents the payment request payload
type ZarinpalRequest struct {
	MerchantID  string `json:"merchant_id"`
	Amount      int    `json:"amount"`
	CallbackURL string `json:"callback_url"`
	Description string `json:"description"`
}

// ZarinpalRequestResponse represents the response from the request
type ZarinpalRequestResponse struct {
	Data struct {
		Code      int    `json:"code"`
		Authority string `json:"authority"`
	} `json:"data"`
	Errors map[string]string `json:"errors"`
}

// RequestPayment sends a payment request to Zarinpal and returns the authority
func RequestPayment(amount int) (string, error) {
	req := ZarinpalRequest{
		MerchantID:  config.AppConfig.ZarinpalMerchantID,
		Amount:      amount,
		CallbackURL: config.AppConfig.CallbackURL,
		Description: "کمک مالی برای خرید گیتار الکتریک امیرحسین",
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post("https://api.zarinpal.com/pg/v4/payment/request.json", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result ZarinpalRequestResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	// Check if request was successful
	if result.Data.Code == 100 {
		return result.Data.Authority, nil
	}
	return "", errors.New("failed to request payment: " + getErrorMessage(result.Errors))
}

// getErrorMessage converts Zarinpal error map to a string
func getErrorMessage(errors map[string]string) string {
	var errorMsg string
	for key, value := range errors {
		errorMsg += key + ": " + value + "; "
	}
	return errorMsg
}
