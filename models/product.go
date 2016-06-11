package models

const (
	apiProductByID = "products/%d"
)

// Product is a model of a single product
// and its characteristics.
type Product struct {
	ID    int
	Face  string
	Size  int
	Price int
}
