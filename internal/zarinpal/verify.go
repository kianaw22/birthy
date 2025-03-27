package zarinpal

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/kianaw22/birthy/config"
)

// ZarinpalVerifyRequest represents the verification request payload
type ZarinpalVerifyRequest struct {
	MerchantID string `json:"merchant_id"`
	Amount     int    `json:"amount"`
	Authority  string `json:"authority"`
}

// ZarinpalVerifyResponse represents the response from the verification
type ZarinpalVerifyResponse struct {
	Data struct {
		Code  int   `json:"code"`
		RefID int64 `json:"ref_id"`
	} `json:"data"`
	Errors map[string]string `json:"errors"`
}

// VerifyPayment verifies the payment status with Zarinpal
func VerifyPayment(authority string, amount int) (bool, int64, error) {
	req := ZarinpalVerifyRequest{
		MerchantID: config.AppConfig.ZarinpalMerchantID,
		Amount:     amount,
		Authority:  authority,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return false, 0, err
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post("https://api.zarinpal.com/pg/v4/payment/verify.json", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, 0, err
	}
	defer resp.Body.Close()

	var result ZarinpalVerifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, 0, err
	}

	// Check if payment verification was successful
	if result.Data.Code == 100 {
		return true, result.Data.RefID, nil
	}
	return false, 0, errors.New("failed to verify payment: " + getErrorMessage(result.Errors))
}
