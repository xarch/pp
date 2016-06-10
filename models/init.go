// Package models provides functions and methods for accessing
// data extracted from the remote API and types for its representation.
package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// remoteAPI is the address of a service that
	// returns information about purchases, users, and products.
	remoteAPI string
)

// get does a request to the specified URI, makes sure the status
// code is 200 OK, and returns the response body or a non-nil error.
func get(uri string) ([]byte, error) {
	// Do a request and make sure it is successfull.
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Make sure the status code is 200 OK.
	if sc := res.StatusCode; sc != http.StatusOK {
		return nil, fmt.Errorf(`unexpected status code "%d"`, sc)
	}

	// Try to read the body of the response.
	// NB: We assume here that the purchases API microservice is trusted
	// and thus wouldn't intentionally return an insanely large response.
	// Otherwise, it would be a good idea to limit the number of bytes
	// we read.
	return ioutil.ReadAll(res.Body)
}

// Init is a functions that initializes the models.
func Init(api string) {
	remoteAPI = api
}
