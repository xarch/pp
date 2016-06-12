// Package handlers provides the standard library compatable HTTP handler
// functions that implement the business logic of the application.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// renderJSON returns a handler function that renders the object it gets
// as an input argument with the specified status code.
// Content-Type of the response is "application/json".
func renderJSON(status int, obj interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Try to marshal the input object.
		res, err := json.Marshal(obj)
		if err != nil {
			// Run the error handler if the object cannot be marshelled.
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Otherwise, render the JSON.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, err = w.Write(res)
		if err != nil {
			log.Println(err)
		}
	}
}

// renderText returns a handler function that renders the text it gets
// as an input argument with the specified status code.
// Content-Type of the response is "plain/text".
func renderText(status int, text string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Render the text.
		w.Header().Set("Content-Type", "plain/text")
		w.WriteHeader(status)
		w.Write([]byte(text))
	}
}
