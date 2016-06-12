package models

import (
	"fmt"
)

const (
	apiProductByID = "products/%d"
)

// productData is a type that represents data
// returned by the product API.
type productData struct {
	Data *Product `json:"product"`
}

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
func ProductByID(id int) (*Product, error) {
	var d productData
	err := objectFromURN(fmt.Sprintf(apiProductByID, id), &d)
	if err != nil {
		return nil, err
	}
	return d.Data, nil
}
