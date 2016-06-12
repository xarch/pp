// Package handlers provides the standard library compatable HTTP handler
// functions that implement the business logic of the application.
package handlers

import (
	"encoding/json"
	"fmt"
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
			renderError(err)
			return
		}

		// Otherwise, render the JSON.
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

// renderText returns a handler function that renders the text it gets
// as an input argument with the specified status code.
// Content-Type of the response is "plain/text".
func renderText(status int, text string, params ...interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "plain/text")
		w.Write([]byte(fmt.Sprintf(text, params...)))
	}
}

// renderError returns a handler function for rendering an internal server error.
func renderError(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf(`Unexpected error: "%s"`, err.Error())
		renderText(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}
