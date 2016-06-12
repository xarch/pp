package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRenderJSON(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	expH1 := "application/json"
	exp1 := `{"a":1,"b":2,"c":3}`

	// Trying to render a valid object.
	obj1 := map[string]int{"a": 1, "b": 2, "c": 3}
	h1 := renderJSON(http.StatusTeapot, obj1)
	h1(w, r)

	if c := http.StatusTeapot; w.Code != c {
		t.Errorf("Expected status %d, got %d.", c, w.Code)
	}
	if s := w.Body.String(); s != exp1 {
		t.Errorf(`Incorrect response. Expected "%s", got "%s".`, exp1, s)
	}
	if h := w.Header().Get("Content-Type"); h != expH1 {
		t.Errorf(`Incorrect content type. Expected "%s", got "%s".`, expH1, h)
	}

	w = httptest.NewRecorder()
	expH2 := "plain/text"
	exp2 := http.StatusText(http.StatusInternalServerError)

	// Making sure error is rendered if there are
	// isues with marshaling.
	obj2 := map[int]int{1: 1, 2: 2, 3: 3}
	h2 := renderJSON(http.StatusTeapot, obj2)
	h2(w, r)

	if c := http.StatusInternalServerError; w.Code != c {
		t.Errorf("Expected status %d, got %d.", c, w.Code)
	}
	if s := w.Body.String(); s != exp2 {
		t.Errorf(`Incorrect response. Expected "%s", got "%s".`, exp2, s)
	}
	if h := w.Header().Get("Content-Type"); h != expH2 {
		t.Errorf(`Incorrect content type. Expected "%s", got "%s".`, expH2, h)
	}
}
