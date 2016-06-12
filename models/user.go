package models

import (
	"fmt"
)

const (
	apiUserByNickname = "users/%s"
)

// User is a type that represents a single customer
// who buyes a product.
type User struct {
	Nickname string `json:"username"`
	Email    string `json:"email"`
}

// UserByNickname gets a username and returns an associated user
// if he/she does exist. It returns a non-nil error otherwise
// or if a connection to the API cannot be established.
func UserByNickname(nickname string) (u *User, err error) {
	err = objectFromURN(fmt.Sprintf(apiUserByNickname, nickname), &u)
	return
}
