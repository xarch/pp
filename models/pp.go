package models

// PopularPurchases model is a convinience for []PopularPurchase
// that provides helper methods.
type PopularPurchases []PopularPurchase

// PopularPurchase is a model that combines information
// about a purchase, product, and recent customers.
type PopularPurchase struct {
	ID      int
	Product *Product
	Recent  []string
}

// PopularPurchasesByNickname returns information about the most popular
// purchases using the following algorithm:
// 1. Fetch N recent purchases for the user.
// 2. For each of that product get a list of people who previously purchased it.
// 3. Fetch the products' details.
// 4. Sort the result: elements with the highest number of purchasers first.
// NB: Operations of the function are not atomic!
/*func PopularPurchasesByNickname(nickname string, limit uint) (PopularPurchases, error) {
	// Make sure the requested user does exist.
	u, err := UserByNickname(nickname)
	if err != nil {
		return nil, err
	}

	// Get recent purchases of the current user.
	ps, err := PurchasesByUsername(u.Nickname,
}*/
