package models

import (
	"encoding/json"
	"net/http"
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
	a.HandleFunc("/purchases/by_user/test", http.NotFoundHandler().ServeHTTP)
	a.HandleFunc("/purchases/by_user/{id}", testPurchasesByArgH)
	a.HandleFunc("/purchases/by_product/{id}", testPurchasesByArgH)

	// Creating a test server with the API.
	s := httptest.NewServer(a)
	defer s.Close()

	// Setting the API's URI.
	Init(s.URL + "/api/")

	// Check the case when the API's response is a valid JSON.
	ps, err := PurchasesByProductID("xxx", 0)
	if err != nil || !reflect.DeepEqual(ps, testPurchases) {
		t.Errorf(`Expected %v, nil. Got %v, "%v".`, testPurchases, ps, err)
	}
	ps, err = PurchasesByUsername("xxx", 0)
	if err != nil || !reflect.DeepEqual(ps, testPurchases) {
		t.Errorf(`Expected %v, nil. Got %v, "%v".`, testPurchases, ps, err)
	}

	// Make sure that there's an error returned if response status is incorrect.
	ps, err = PurchasesByUsername("test", 0)
	if ps != nil || err == nil {
		t.Errorf(`Expected no result and an error. Got %v, "%v".`, ps, err)
	}
}

// testPurchasesByArgH is a handler that imitates the third
// party API that provides purchases by username / id.
var testPurchasesByArgH = func(w http.ResponseWriter, r *http.Request) {
	// Transform the test list of purchases into JSON.
	res, err := json.Marshal(testPurchases)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Render the result.
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

var testPurchases = []Purchase{
	{123, 321, "JohnDoe", time.Now()},
	{222, 444, "Mr.X", time.Now()},
}
