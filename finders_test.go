package oomg

import (
	"github.com/ajoyac/oomg/testhelper"
	"testing"
)

func TestDbAll(t *testing.T) {

	_ = Connect(ConnectOptions{URL: "mongodb://localhost:27017", DbName: "foo"})
	t.Run("Find All empty", func(t *testing.T) {

		foo := make([]testhelper.Foo, 0)
		err := All(&foo)
		if err != nil {
			t.Error(err)
			return
		}
	})

}

func TestDbOne(t *testing.T) {

	_ = Connect(ConnectOptions{URL: "mongodb://localhost:27017", DbName: "foo"})
	t.Run("Find One ", func(t *testing.T) {
		foo := testhelper.Foo{}
		err := One(&foo)
		if err != nil {
			t.Error(err)
			return
		}
	})

}
