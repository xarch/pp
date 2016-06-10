package models

import (
	"encoding/json"
	"fmt"
	"path"
	"time"
)

var (
	apiPurchasesByUsername = path.Join(remoteAPI, "purchases/by_user/%s?limit=%d")
	apiPurchasesByProduct  = path.Join(remoteAPI, "purchases/by_product/%s?limit=%d")
)

// Purchase type represents information a single order.
type Purchase struct {
	ID        int       `json:"id,omitempty"`
	ProductID int       `json:"product_id"`
	Username  string    `json:"username"`
	Date      time.Time `json:"date,omitempty"`
}

// PurchasesByUsername gets a username and a limit number as input
// arguments and returns a list of purchases that belong to the
// related user.
func PurchasesByUsername(username string, limit uint) ([]Purchase, error) {
	ps := []Purchase{}

	// Do a GET request to the remote server.
	res, err := get(fmt.Sprintf(apiPurchasesByUsername, username, limit))
	if err != nil {
		return nil, err
	}

	// Try to parse the body of the response.
	err = json.Unmarshal(res, &ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}
