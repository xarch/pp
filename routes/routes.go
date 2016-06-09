// Package routes defines the way this application will
// route user requests.
package routes

import (
	"net/http"
)

// List is a slice of routes this application must handle.
var List = []struct {
	// Methods is a full list of HTTP methods that must be supported
	// by a route. A slice is used instead of map as the complexity of the
	// former is O(n) and of the latter is O(1+c). And if n is small,
	// O(n) < O(1+c).
	Methods []string

	Pattern string
	Handler http.Handler
}{}
