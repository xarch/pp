package models

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	apiPurchasesByUsername = "purchases/by_user/%s?limit=%d"
	apiPurchasesByProduct  = "purchases/by_product/%s?limit=%d"
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
	return purchasesFromURN(fmt.Sprintf(apiPurchasesByUsername, username, limit))
}

// PurchasesByProductID gets an ID of a product and a limit number as
// input arguments and returns a list of all purchases of the product.
func PurchasesByProductID(id string, limit uint) ([]Purchase, error) {
	return purchasesFromURN(fmt.Sprintf(apiPurchasesByProduct, id, limit))
}

// purchasesFromURN gets a URN of an API, makes a GET request to get the response,
// parses it, and returns an unmarshalled result.
// If something goes wrong, an error is returned as a second argument.
func purchasesFromURN(urn string) ([]Purchase, error) {
	ps := []Purchase{}

	// Do a GET request to the remote server.
	res, err := get(remoteAPI + urn)
	if err != nil {
		return nil, err
	}

	// Try to unmarshal the body of the response.
	err = json.Unmarshal(res, &ps)
	if err != nil {
		return nil, err
	}

	return ps, nil
}
