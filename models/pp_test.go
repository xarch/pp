package models

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"
)

func TestPopularPurchaseUnmarshalJSON(t *testing.T) {
	src := []byte(`{"id":123,"face":"xxx","size":15,"price":21,"recent":["John.Doe","Jane.Roe","James_Smith"]}`)
	exp := PopularPurchase{
		ID: 123,
		Product: &Product{
			ID:    0,
			Face:  "xxx",
			Size:  15,
			Price: 21,
		},
		Recent: []string{"John.Doe", "Jane.Roe", "James_Smith"},
	}
	var pp PopularPurchase
	err := json.Unmarshal(src, &pp)
	if err != nil {
		t.Error(err)
	}
	if pp.ID != exp.ID || !reflect.DeepEqual(pp.Recent, exp.Recent) ||
		!reflect.DeepEqual(*pp.Product, *exp.Product) {

		t.Errorf(`Incorrect result. Expected: %#v (%v). Got: %#v (%v).`, exp, exp.Product, pp, pp.Product)
	}
}

func TestPopularPurchaseMarshalJSON(t *testing.T) {
	p := PopularPurchase{
		ID: 123,
		Product: &Product{
			ID:    321,
			Face:  "xxx",
			Size:  15,
			Price: 21,
		},
		Recent: []string{"John.Doe", "Jane.Roe", "James_Smith"},
	}
	exp := `{"id":123,"face":"xxx","size":15,"price":21,"recent":["John.Doe","Jane.Roe","James_Smith"]}`
	res, err := p.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if reflect.DeepEqual(res, []byte(exp)) {
		t.Errorf(`Incorrect JSON returned. Expected "%v". Got "%v".`, exp, string(res))
	}
}

func TestPopularPurchases_Sorting(t *testing.T) {
	ps := PopularPurchases{
		{Recent: []string{"a", "b", "c", "d", "e", "f"}},
		{Recent: []string{"a", "b"}},
		{Recent: []string{"a", "b", "c"}},
		{Recent: []string{"a", "b", "c", "d", "e", "f"}},
		{Recent: []string{"a", "b", "c", "d"}},
		{Recent: []string{"a"}},
	}
	exp := PopularPurchases{
		{Recent: []string{"a", "b", "c", "d", "e", "f"}},
		{Recent: []string{"a", "b", "c", "d", "e", "f"}},
		{Recent: []string{"a", "b", "c", "d"}},
		{Recent: []string{"a", "b", "c"}},
		{Recent: []string{"a", "b"}},
		{Recent: []string{"a"}},
	}
	sort.Sort(ps)
	if !reflect.DeepEqual(ps, exp) {
		t.Errorf("Result is not sorted correctly. Expected %v. Got %v.", exp, ps)
	}
}
