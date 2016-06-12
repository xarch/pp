package handlers

import (
	"flag"
	"net/http"
	"sort"
	"time"

	"github.com/alkchr/pp/models"

	"github.com/allegro/bigcache"
)

const (
	notFoundMsg = "User with username of '%s' was not found"
)

var (
	apiURI               = flag.String("purchases:api.uri", "http://127.0.0.1/", "Address of the purchases API")
	recentPurchasesNum   = flag.Uint("purchases:recent.num", 5, "Number of recent purchases to use")
	purchasersPerProduct = flag.Uint("purchases:customers.limit", 150, "Maximum number of customers per product")

	expiratSecs = flag.Int64("cache:expiration.seconds", 300, "For how many seconds response must be cached")
)

// PopularPurchases returns a handler function that uses the following algorithm
// for rendering the list of the most popular purchases:
// 1. Fetch N recent purchases for the user.
// 2. For each of that product get a list of people who previously purchased it.
// 3. Fetch the products' details.
// 4. Sort the result: elements with the highest number of purchasers first.
// NB: A primitive version of cache is implemented. It does exactly one thing:
// caches the final result for N seconds.
func PopularPurchases(c map[string]string) http.HandlerFunc {
	// NB: all input arguments must be validated using
	// a whitelist of allowed characters. We are not doing that for the username as
	// router already handles this for us.
	nickname := c["username"]

	// Check whether we have already had a cached version of result
	// for the requested user.
	if bs, err := cache.Get(nickname); err == nil {
		return render(http.StatusNotModified, bs, "application/json")
	}

	//
	// NB: Operations are not atomic.
	//

	// Make sure the requested user does exist.
	// If he/she doesn't render a not found error.
	u, err := models.UserByNickname(nickname)
	if err != nil {
		return renderText(http.StatusNotFound, notFoundMsg, nickname)
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

	// Sort the result, add it to the cache, and render it.
	sort.Sort(pp)
	return renderJSON(http.StatusOK, pp, func(d []byte) {
		cache.Set(nickname, d)
	})
}

// Init is a function that should be used for initialization
// of handlers.
func Init() {
	// Initialize the models and prepare for connection.
	models.Init(*apiURI)

	// Initialize cache with the default parameters.
	cache, _ = bigcache.NewBigCache(
		bigcache.DefaultConfig(time.Duration(*expiratSecs) * time.Second),
	)
}
