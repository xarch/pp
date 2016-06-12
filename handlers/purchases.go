package handlers

import (
	"net/http"
)

// PopularPurchases is a handler function that uses the following algorithm
// for rendering the list of the most popular purchases:
// 1. Fetch 5 recent purchases for the user.
// 2. For each of that product get a list of people who previously purchased it.
// 3. Fetch the products' details.
// 4. Sort the result: elements with the highest number of purchasers first.
func PopularPurchases(w http.ResponseWriter, r *http.Request, c map[string]string) {
	// NB: all input arguments must be validated using
	// a whitelist of allowed characters. We are not doing that for the username as
	// router already handles this for us.
	uname := r.FormValue("username")

	renderJSON(http.StatusOK, uname)
}
