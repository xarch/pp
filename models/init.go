// Package models provides functions and methods for accessing
// data extracted from the remote API and types for its representation.
package models

var (
	// remoteAPI is the address of a service that
	// returns information about purchases, users, and products.
	remoteAPI string
)

// Init is a functions that initializes the models.
func Init(api string) {
	remoteAPI = api
}
