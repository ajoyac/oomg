package oomg

import (
	"github.com/ajoyac/oomg/testhelper"
	"testing"
)

func TestDbAll(t *testing.T) {

	db := DB()
	foo := make([]testhelper.Foo, 0)
	err := db.All(&foo)
	if err != nil {
		t.Error(err)
		return
	}
	if len(foo) != 0 {
		t.Errorf("Expected empty found %v elements", len(foo))
	}

}
