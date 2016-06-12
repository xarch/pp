package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/alkchr/pp/models"

	"github.com/gorilla/mux"
)

func TestPopularPurchases(t *testing.T) {
	m := mux.NewRouter()
	a := m.Path("/api/").Subrouter()

	a.HandleFunc("/purchases/by_user/testUser", http.NotFoundHandler().ServeHTTP)
	a.HandleFunc("/purchases/by_product/666", http.NotFoundHandler().ServeHTTP)

	a.HandleFunc("/products/{id}", productByIDH)
	a.HandleFunc("/purchases/by_user/{username}", purchasesByUsernameH)
	a.HandleFunc("/purchases/by_product/{id}", purchasesByProductIDH)
	a.HandleFunc("/users/{username}", userByNicknameH)

	s := httptest.NewServer(a)
	defer s.Close()

	*apiURI = s.URL + "/api/"
	Init()

	r, _ := http.NewRequest("GET", "/", nil)

	for i, v := range []struct {
		username                string
		expStatus               int
		expBody, expContentType string
	}{
		{"doesnt_exist", http.StatusNotFound, fmt.Sprintf(notFoundMsg, "doesnt_exist"), "plain/text"},
		{"testUser", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "plain/text"},
		{"testUser666", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), "plain/text"},
	} {
		w := httptest.NewRecorder()
		PopularPurchases(map[string]string{"username": v.username})(w, r)
		if w.Code != v.expStatus {
			t.Errorf("Test %d: Expected status %d, got %d.", i, v.expStatus, w.Code)
		}
		if ct := w.Header().Get("Content-Type"); ct != v.expContentType {
			t.Errorf(`Test %d: Incorrect content type. Expected "%s", got "%s".`, i, v.expContentType, ct)
		}
		if w.Body.String() != v.expBody {
			t.Errorf(`Test %d: Incorrect response body. Expected "%s", got "%s".`, i, v.expBody, w.Body.String())
		}
	}
}

// Test API Server.
// Another alternative would be to use integration tests
// and a daw-purchases service. But I've decided not to bring
// nodejs as a dependency for such a small project.

// productByIDH handler implements "GET /api/products/{id:[0-9]}
var productByIDH = func(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		renderError(err)
		return
	}
	for i := range testData.Products {
		if testData.Products[i].ID == int(id) {
			renderJSON(http.StatusOK, map[string]interface{}{
				"product": testData.Products[i],
			})(w, r)
			return
		}
	}
	renderJSON(http.StatusOK, map[string]interface{}{})(w, r)
}

// purchasesByUsernameH handler implements "GET /api/purchases/by_user/{username:[A-Za-z0-9_.-]}
var purchasesByUsernameH = func(w http.ResponseWriter, r *http.Request) {
	ps := models.Purchases{}
	for i := range testData.Purchases {
		if testData.Purchases[i].Username != mux.Vars(r)["username"] {
			continue
		}
		ps = append(ps, testData.Purchases[i])
	}
	renderJSON(http.StatusOK, map[string]interface{}{
		"purchases": ps,
	})(w, r)
}

// purchasesByProductIDH handler implements "GET /api/purchases/by_product/{id:[0-9]}
var purchasesByProductIDH = func(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		renderError(err)(w, r)
		return
	}
	ps := models.Purchases{}
	for i := range testData.Purchases {
		if testData.Purchases[i].ProductID != int(id) {
			continue
		}
		ps = append(ps, testData.Purchases[i])
	}
	renderJSON(http.StatusOK, map[string]interface{}{
		"purchases": ps,
	})(w, r)
}

// userByNicknameH handler implements "GET /api/users/{nickname:[A-Za-z0-9_.-]}
var userByNicknameH = func(w http.ResponseWriter, r *http.Request) {
	for i := range testData.Users {
		if testData.Users[i].Nickname == mux.Vars(r)["username"] {
			renderJSON(http.StatusOK, map[string]interface{}{
				"user": testData.Users[i],
			})(w, r)
			return
		}
	}
	renderJSON(http.StatusOK, map[string]interface{}{})(w, r)
}

var testData data

type data struct {
	Purchases models.Purchases `json:"purchases"`
	Products  []models.Product `json:"products"`
	Users     []models.User    `json:"users"`
}

func init() {
	d, err := ioutil.ReadFile("./testdata/data.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(d, &testData)
	if err != nil {
		panic(err)
	}

	// Intentionally inserting inconsistent data for testing.
	testData.Users = append(testData.Users, models.User{
		Nickname: "testUser",
		Email:    "test@user",
	})
	testData.Users = append(testData.Users, models.User{
		Nickname: "testUser666",
		Email:    "test@user666",
	})
	testData.Purchases = append(testData.Purchases, models.Purchase{
		Username:  "testUser666",
		ProductID: 666,
	})
	testData.Products = append(testData.Products, models.Product{
		ID: 666,
	})
}
