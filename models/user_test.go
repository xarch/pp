package models

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestUserByUsername(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/users/{id}", renderJSONHandlerFn(testUser))

	// Creating a test server imitating a third party API.
	s := httptest.NewServer(r)
	defer s.Close()

	// Setting the API's URI.
	Init(s.URL + "/api/")

	// Make sure ProductByID returns a valid Product.
	u, err := UserByNickname("some.username")
	if err != nil || reflect.DeepEqual(testUser, u) {
		t.Errorf(`Expected %v, "nil". Got %v, "%v".`, testProduct, u, err)
	}
}

var testUser = User{
	Nickname: "some.username",
	Email:    "some@username.com",
}
