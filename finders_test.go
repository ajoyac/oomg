package oomg

import (
	"github.com/ajoyac/oomg/testhelper"
	"testing"
)

func TestDbAll(t *testing.T) {

	_ = Connect(ConnectOptions{URL: "mongodb://localhost:27017", DbName: "foo"})

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
