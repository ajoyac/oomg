package oomg

import (
	"github.com/ajoyac/oomg/testhelper"
	"testing"
)

func TestDbCreate(t *testing.T) {

	_ = Connect(ConnectOptions{URL: "mongodb://localhost:27017", DbName: "foo"})

	t.Run("Save One ", func(t *testing.T) {
		db := DB()
		foo := testhelper.Foo{}
		err := db.Create(&foo)
		if err != nil {
			t.Error(err)
			return
		}
	})

}
