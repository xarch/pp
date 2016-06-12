package models

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestObjectFromURN(t *testing.T) {
	var obj interface{}
	err := objectFromURN("urn_that_doesnt_exist", &obj)
	if obj != nil || err == nil {
		t.Errorf(`Expected no result and an error. Got %v, "%v".`, obj, err)
	}

	// Imitating a third party API.
	r := mux.NewRouter()
	r.HandleFunc("/incorrect/status", http.NotFoundHandler().ServeHTTP)
	r.HandleFunc("/non/json", emptyH)

	// Creating a test server with the API.
	s := httptest.NewServer(r)
	defer s.Close()

	// Setting the API's URI.
	Init(s.URL + "/api/")

	// Checking all kinds of errors:
	// 1. Incorrect status code of response.
	// 2. Non-json response.
	for _, urn := range []string{"/incorrect/status", "/non/json"} {
		var obj interface{}
		err := objectFromURN(remoteAPI+urn, &obj)
		if err == nil || obj != nil {
			t.Errorf(`Expected nil, error. Got %v, %v.`, obj, err)
		}
	}
}

// renderJSONHandlerFn gets an object and returns a handler
// function that will render it along with using application/json
// content-type header and 200 Status OK.
func renderJSONHandlerFn(obj interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Transform the test product into JSON.
		res, err := json.Marshal(obj)
		if err != nil {
			panic(err)
		}

		// Render the result.
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

// emptyH is a handler that renders an empty page.
var emptyH = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
