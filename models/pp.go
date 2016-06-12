package models

import (
	"encoding/json"
)

// ppData is the way popular purchases are expected to be
// represented in the response of the application.
type ppData struct {
	ID     int      `json:"id"`
	Face   string   `json:"face"`
	Price  int      `json:"price"`
	Size   int      `json:"size"`
	Recent []string `json:"recent"`
}

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

// UnmarshalJSON implements json.Unmarshaler interface to unmarshal
// JSON data into the object.
// It is used for testing purposes only.
func (m *PopularPurchase) UnmarshalJSON(d []byte) error {
	var obj ppData
	err := json.Unmarshal(d, &obj)
	*m = PopularPurchase{
		ID: obj.ID,
		Product: &Product{
			Face:  obj.Face,
			Price: obj.Price,
			Size:  obj.Size,
		},
		Recent: obj.Recent,
	}
	return err
}

// MarshalJSON implements json.Marshaler interface to marshal
// the object into valid JSON.
func (m *PopularPurchase) MarshalJSON() ([]byte, error) {
	return json.Marshal(ppData{
		ID:     m.ID,
		Face:   m.Product.Face,
		Price:  m.Product.Price,
		Size:   m.Product.Size,
		Recent: m.Recent,
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
