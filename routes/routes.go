// Package routes defines the way pp application will
// route user requests.
// Use it as follows:
//	s := http.Server{
//		...
//		Handler: routes.Handler(),
//	}
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Handler returns an HTTP handler that multiplexes requests
// to the application's handlers.
func Handler() http.Handler {
	// Allocate a new router. Gorilla router with O(n) complexity is used as there
	// is just one route. Replace it by a trie based multiplexer if the number
	// of routes is growing.
	r := mux.NewRouter()

	// TODO: for type safety use http.Method{Name} constants instead if manually writing
	// method names when Go 1.7 is stable and no support of other versions is required.
	api := r.Path("/api/").Subrouter()
	api.Methods(
		"GET",
	).Path("/recent_purchases/{username:[A-Za-z0-9-_.]+}").Handler(http.NotFoundHandler())

	return r
}
