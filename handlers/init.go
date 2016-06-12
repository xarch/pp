// Package handlers provides the standard library compatible HTTP handler
// functions that implement the business logic of the application.
package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/allegro/bigcache"
)

// cache is a fast and concurrent in-memory store.
// If multiple instances of the application are started,
// no state of cache will be shared.
// Use memcached, redis or groupcache if that's required.
var cache *bigcache.BigCache

// renderJSON returns a handler function that renders the object it gets
// as an input argument with the specified status code.
// If the object's marshaling is successfull, an input function is executed.
// Content-Type of the response is "application/json".
func renderJSON(status int, obj interface{}, fn func([]byte)) http.HandlerFunc {
	// Try to marshal the input object.
	res, err := json.Marshal(obj)
	if err != nil {
		// Run the error handler if the object cannot be marshelled.
		return renderError(err)
	}
	fn(res) // Successfully marshalled the result, start the function.
	return render(status, res, "application/json")
}

// renderText returns a handler function that renders the text it gets
// as an input argument with the specified status code.
// Content-Type of the response is "plain/text".
func renderText(status int, text string, params ...interface{}) http.HandlerFunc {
	return render(status, []byte(fmt.Sprintf(text, params...)), "plain/text")
}

// renderError returns a handler function for rendering an internal server error.
func renderError(err error) http.HandlerFunc {
	log.Printf(`Unexpected error: "%s"`, err.Error())
	return renderText(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

// render returns a handler for rendering arbitrary data with the specified
// status code and content type.
func render(status int, data []byte, ct string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Header().Set("Content-Type", ct)
		w.Write(data)
	}
}
