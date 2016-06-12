package models

import (
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestPurchasesByArgument(t *testing.T) {
	// Imitating a third party API. Unlike the real one this always
	// returns the same JSON response.
	r := mux.NewRouter()
	a := r.Path("/api/").Subrouter()
	a.HandleFunc("/purchases/by_user/incorrect", emptyH)
	a.HandleFunc("/purchases/by_user/{id}", renderJSONHandlerFn(testPurchases))
	a.HandleFunc("/purchases/by_product/{id}", renderJSONHandlerFn(testPurchases))

	// Creating a test server with the API.
	s := httptest.NewServer(a)
	defer s.Close()

	// Setting the API's URI.
	Init(s.URL + "/api/")

	// Check the case when the API's response is a valid JSON.
	ps, err := PurchasesByUsername("xxx", 0)
	if err != nil || !deepEqualPurchases(testPurchases, ps) {
		t.Errorf(`Expected %v, "nil". Got %v, "%v".`, testPurchases, ps, err)
	}
	ps, err = PurchasesByProductID(999, 0)
	if err != nil || !deepEqualPurchases(testPurchases, ps) {
		t.Errorf(`Expected %v, "nil". Got %v, "%v".`, testPurchases, ps, err)
	}

	// Check non-existent user.
	ps, err = PurchasesByUsername("incorrect", 0)
	if ps != nil || err == nil {
		t.Errorf(`Expected no result and an error. Got %v, "%v".`, ps, err)
	}
}

func TestPurchasePopular_FailToGetProduct(t *testing.T) {
	p := Purchase{ProductID: 123}
	pp, err := p.Popular(1)
	if pp != nil || err == nil {
		t.Errorf("Cannot get product. Expected: nil, error. Got: %v, %v.", pp, err)
	}
}

func TestPurchasesPopular(t *testing.T) {
	// Imitating a third party API. Unlike the real one this always
	// returns the same JSON response.
	r := mux.NewRouter()
	a := r.Path("/api/").Subrouter()
	a.HandleFunc("/products/{id}", renderJSONHandlerFn(testProduct))
	a.HandleFunc("/purchases/by_product/777", emptyH)
	a.HandleFunc("/purchases/by_product/{id}", renderJSONHandlerFn(testPurchases))

	// Creating a test server with the API.
	s := httptest.NewServer(a)
	defer s.Close()

	// Setting the API's URI.
	Init(s.URL + "/api/")

	// If purchases cannot be fetched, an error is expected.
	ps := Purchases{{ProductID: 777}}
	pp, err := ps.Popular(1)
	if pp != nil || err == nil {
		t.Errorf("Cannot get purchases. Expected: nil, error. Got: %v, %v.", pp, err)
	}

	// Checking a case when no error is expected.
	ps = Purchases{{ID: 456, ProductID: testProduct.ID}}
	pp, err = ps.Popular(1)
	exp := PopularPurchases{
		{ID: 456, Product: &testProduct, Recent: []string{"JohnDoe", "Mr.X"}},
	}
	if err != nil || !reflect.DeepEqual(pp, exp) {
		t.Errorf(`Incorrect popular purchases. Expected %v, "nil". Got %v, "%v".`, exp, pp, err)
	}
}

func TestPurchasesCustomerUsernames(t *testing.T) {
	ps := Purchases{
		{Username: "aaa"},
		{Username: "bbb"},
		{Username: "ccc"},
		{Username: "bbb"},
		{Username: "bbb"},
		{Username: "aaa"},
		{Username: "aaa"},
		{Username: "ddd"},
	}
	exp := []string{"aaa", "bbb", "ccc", "ddd"}
	if us := ps.CustomerUsernames(); !reflect.DeepEqual(us, exp) {
		t.Errorf("Incorrect usernames. Expected: %v, got: %v.", exp, us)
	}
}

// deepEqualPurchases compares two sets of purchases and makes sure they are equal
// to each other. reflect.Deep is not an option as Location field
// of Date parameter is a pointer that will be different in the original
// and marshalled objects.
func deepEqualPurchases(ps1, ps2 []Purchase) bool {
	if len(ps1) != len(ps2) {
		return false
	}
	for i := 0; i < len(ps1); i++ {
		if ps1[i].ID != ps2[i].ID || ps1[i].ProductID != ps2[i].ProductID ||
			ps1[i].Username != ps2[i].Username || ps1[i].Date.String() != ps2[i].Date.String() {

			return false
		}
	}
	return true
}

var testPurchases = []Purchase{
	{123, 321, "JohnDoe", time.Now().Local().Round(time.Minute)},
	{222, 444, "Mr.X", time.Now().Local().Round(time.Minute)},
}
