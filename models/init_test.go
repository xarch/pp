package models

import (
	"testing"
)

func TestObjectFromURN(t *testing.T) {
	var obj interface{}
	err := objectFromURN("urn_that_doesnt_exist", &obj)
	if obj != nil || err == nil {
		t.Errorf(`Expected no result and an error. Got %v, "%v".`, obj, err)
	}
}
