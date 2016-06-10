package models

import (
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
// If limit is equal to 0, it assumes there is no limit.
func PurchasesByUsername(username string, limit uint) []Purchase {
	return nil
}
