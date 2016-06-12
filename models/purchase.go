package models

import (
	"fmt"
	"time"
)

const (
	apiPurchasesByUsername = "purchases/by_user/%s?limit=%d"
	apiPurchasesByProduct  = "purchases/by_product/%d?limit=%d"
)

// Purchases models contains information about
// multiple orders of a product / products.
type Purchases []Purchase

// purchasesData is a type that's used for representation of
// Purchases by the third party API.
type purchasesData struct {
	Data Purchases `json:"purchases"`
}

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
// NB: According to the gdocs the API must return the most recent purchases
// (i.e. they are sorted by the date).
// Service "daw-purchases" by-default returns a random unsorted data instead.
func PurchasesByUsername(username string, limit uint) (Purchases, error) {
	// Get the user's recent purchases.
	var d purchasesData
	err := objectFromURN(fmt.Sprintf(apiPurchasesByUsername, username, limit), &d)
	if err != nil {
		return nil, err
	}
	return d.Data, nil
}

// PurchasesByProductID gets an ID of a product and a limit number as
// input arguments and returns a list of all purchases of the product.
func PurchasesByProductID(id int, limit uint) (Purchases, error) {
	var d purchasesData
	err := objectFromURN(fmt.Sprintf(apiPurchasesByProduct, id, limit), &d)
	if err != nil {
		return nil, err
	}
	return d.Data, nil
}

// CustomerUsernames returns a slice of usernames the purchases' customers.
// Every username is unique, there are no repetitions.
func (m Purchases) CustomerUsernames() (us []string) {
	check := map[string]bool{}
	for i := range m {
		// Ignore the username we have already added to the list.
		if check[m[i].Username] {
			continue
		}
		us = append(us, m[i].Username)
		check[m[i].Username] = true
	}
	return
}

// Popular allocates and returns a new PopularPurchase.
// NB: This function is not atomic.
func (m Purchase) Popular(limit uint) (*PopularPurchase, error) {
	// Try to get product info.
	p, err := ProductByID(m.ProductID)
	if err != nil {
		return nil, err
	}
	// Try to get buyers of the product (aka customers).
	cs, err := PurchasesByProductID(m.ProductID, limit)
	if err != nil {
		return nil, err
	}
	return &PopularPurchase{
		ID:      m.ID,
		Product: p,
		Recent:  cs.CustomerUsernames(),
	}, nil
}

// Popular allocates and returns new PopularPurchases.
// NB: This function is not atomic.
func (m Purchases) Popular(limit uint) (ps PopularPurchases, err error) {
	for i := range m {
		p, err := m[i].Popular(limit)
		if err != nil {
			return nil, err
		}
		ps = append(ps, *p)
	}
	return
}
