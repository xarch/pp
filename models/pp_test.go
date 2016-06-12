package models

import (
	"reflect"
	"sort"
	"testing"
)

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
