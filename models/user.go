package models

import (
	"path"
)

var (
	apiUserByUsername = path.Join(remoteAPI, "users/%s")
)

// User is a type that represents a single customer
// who buyes a product.
type User struct {
	Nickname string `json:"username"`
	Email    string `json:"email"`
}
