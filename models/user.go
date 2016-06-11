package models

const (
	apiUserByUsername = "users/%s"
)

// User is a type that represents a single customer
// who buyes a product.
type User struct {
	Nickname string `json:"username"`
	Email    string `json:"email"`
}
