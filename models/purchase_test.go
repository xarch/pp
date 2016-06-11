package models

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestPurchasesByArgument(t *testing.T) {
	// Imitating a third party API. Unlike the real one this always
	// returns the same JSON response.
	r := mux.NewRouter()
	a := r.Path("/api/").Subrouter()
	a.HandleFunc("/purchases/by_user/test", http.NotFoundHandler().ServeHTTP)
	a.HandleFunc("/purchases/by_user/empty", emptyH)
	a.HandleFunc("/purchases/by_user/{id}", renderJSONHandlerFn(testPurchases))
	a.HandleFunc("/purchases/by_product/{id}", renderJSONHandlerFn(testPurchases))

	// Creating a test server with the API.
	s := httptest.NewServer(a)
	defer s.Close()

	// Setting the API's URI.
	Init(s.URL + "/api/")

	// Check the case when the API's response is a valid JSON.
	for i, fn := range []func(string, uint) ([]Purchase, error){
		PurchasesByUsername,
		PurchasesByProductID,
	} {
		ps, err := fn("xxx", 0)
		if err != nil || !deepEqualPurchases(testPurchases, ps) {
			t.Errorf(`Test %d: Expected %#v, "nil". Got %#v, "%v".`, i, testPurchases, ps, err)
		}
	}

	// Check all possible errors, including:
	// 1. incorrect response status.
	// 2. invalid JSON.
	for i, arg := range []string{"test", "empty"} {
		ps, err := PurchasesByUsername(arg, 0)
		if ps != nil || err == nil {
			t.Errorf(`Test %d: Expected no result and an error. Got %v, "%v".`, i, ps, err)
		}
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
