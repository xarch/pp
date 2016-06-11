package models

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestObjectFromURN(t *testing.T) {
	var obj interface{}
	err := objectFromURN("urn_that_doesnt_exist", &obj)
	if obj != nil || err == nil {
		t.Errorf(`Expected no result and an error. Got %v, "%v".`, obj, err)
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
