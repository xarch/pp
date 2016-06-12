package models

import (
	"encoding/json"
)

// PopularPurchases model is a convinience for []PopularPurchase
// that provides helper methods.
type PopularPurchases []PopularPurchase

// PopularPurchase is a model that combines information
// about a purchase, product, and recent customers.
type PopularPurchase struct {
	ID      int
	Product *Product
	Recent  []string
}

// MarshalJSON implements json.Marshaler interface to marshal
// the object into valid JSON.
func (m *PopularPurchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(data{
		"id":     m.ID,
		"face":   m.Product.Face,
		"price":  m.Product.Price,
		"size":   m.Product.Size,
		"recent": m.Recent,
	})
}

// Len is part of sort.Interface. It returns a number of elements.
func (m PopularPurchases) Len() int {
	return len(m)
}

// Swap is part of sort.Interface. It gets two indexes and swaps
// related elements.
func (m PopularPurchases) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// Less is part of sort.Interface. It gets two indexes and checks
// whether i-th element is less than j-th one.
// The criteria of comparing is the number of recent buyers.
func (m PopularPurchases) Less(i, j int) bool {
	if len(m[i].Recent) < len(m[j].Recent) {
		return false
	}
	return true
}
