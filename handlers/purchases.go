package handlers

import (
	"flag"
	"net/http"
	"sort"

	"github.com/alkchr/pp/models"
)

var (
	apiURI               = flag.String("purchases:api.uri", "http://127.0.0.1/", "Address of the purchases API")
	recentPurchasesNum   = flag.Uint("purchases:recent.num", 5, "Number of recent purchases to use")
	purchasersPerProduct = flag.Uint("purchases:customers.limit", 150, "Maximum number of customers per product")
)

// PopularPurchases is a handler function that uses the following algorithm
// for rendering the list of the most popular purchases:
// 1. Fetch N recent purchases for the user.
// 2. For each of that product get a list of people who previously purchased it.
// 3. Fetch the products' details.
// 4. Sort the result: elements with the highest number of purchasers first.
// TODO: implement caching.
func PopularPurchases(w http.ResponseWriter, r *http.Request, c map[string]string) http.HandlerFunc {
	// NB: all input arguments must be validated using
	// a whitelist of allowed characters. We are not doing that for the username as
	// router already handles this for us.
	nickname := c["username"]

	//
	// NB: Operations are not atomic.
	//

	// Make sure the requested user does exist.
	// If he/she doesn't render a not found error.
	u, err := models.UserByNickname(nickname)
	if err != nil {
		return renderText(http.StatusNotFound, "User with username of '%s' was not found", nickname)
	}

	// Get current user's recent purchases.
	ps, err := models.PurchasesByUsername(u.Nickname, *recentPurchasesNum)
	if err != nil {
		return renderError(err)
	}

	// Fetch the related information such as recent buyers and product details.
	pp, err := ps.Popular(*purchasersPerProduct)
	if err != nil {
		return renderError(err)
	}

	// Sort the result and render it.
	sort.Sort(pp)
	return renderJSON(http.StatusOK, pp)
}

// Init is a function that should be used for initialization
// of handlers.
func Init() {
	models.Init(*apiURI)
}
