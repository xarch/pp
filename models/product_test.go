package models

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestProductByID(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/products/{id}", renderJSONHandlerFn(testProduct))

	// Creating a test server imitating a third party API.
	s := httptest.NewServer(r)
	defer s.Close()

	// Setting the API's URI.
	Init(s.URL + "/api/")

	// Make sure ProductByID returns a valid Product.
	p, err := ProductByID(123)
	if err != nil || reflect.DeepEqual(testProduct, p) {
		t.Errorf(`Expected %v, "nil". Got %v, "%v".`, testProduct, p, err)
	}
}

var testProduct = productData{
	Data: &Product{
		ID:    123,
		Face:  "xxx",
		Size:  55,
		Price: 99,
	},
}
