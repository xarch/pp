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

	"github.com/alkchr/pp/handlers"

	"github.com/gorilla/mux"
)

// Handler returns an HTTP handler that multiplexes requests
// to the application's handlers.
func Handler() http.Handler {
	// Allocate a new router. Gorilla router with O(n) complexity is used as there
	// is just one route. Replace it by a trie based multiplexer if the number
	// of routes is growing.
	r := mux.NewRouter()
	api := r.Path("/api/").Subrouter()

	// TODO: for type safety use http.Method{Name} constants instead if manually writing
	// method names when Go 1.7 is stable and no support of other versions is required.
	api.HandleFunc(
		"/recent_purchases/{username:[A-Za-z0-9_.-]+}", wrap(handlers.PopularPurchases),
	).Methods("GET")

	return api
}

// wrap is a helper that gets a handler function with the
// third context parameter as input and returns a standard handler function.
// It is used for passing mux's parameters to the handlers.
//
// NB: complexity of getting a single element from a map is O(1+c).
// In comparison, in case of a slice it would be O(n).
// But if n is small, O(n) < O(1+c). Thus, consider
// replacing the context's type if another router is in use.
func wrap(fn func(http.ResponseWriter, *http.Request, map[string]string) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, mux.Vars(r))(w, r)
	}
}
