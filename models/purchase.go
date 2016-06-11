package models

import (
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
func PurchasesByUsername(username string, limit uint) (ps []Purchase, err error) {
	err = objectFromURN(fmt.Sprintf(apiPurchasesByUsername, username, limit), &ps)
	return
}

// PurchasesByProductID gets an ID of a product and a limit number as
// input arguments and returns a list of all purchases of the product.
func PurchasesByProductID(id string, limit uint) (ps []Purchase, err error) {
	err = objectFromURN(fmt.Sprintf(apiPurchasesByProduct, id, limit), &ps)
	return
}
