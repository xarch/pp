package models

import (
	"fmt"
)

const (
	apiProductByID = "products/%d"
)

// Product is a model of a single product
// and its characteristics.
type Product struct {
	ID    int    `json:"id"`
	Face  string `json:"face"`
	Size  int    `json:"size"`
	Price int    `json:"price"`
}

// ProductByID receives a product info with the requested ID
// from a third party API.
func ProductByID(id int) (p *Product, err error) {
	err = objectFromURN(fmt.Sprintf(apiProductByID, id), &p)
	return
}
