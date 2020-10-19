package oomg

import (
	"github.com/ajoyac/oomg/testhelper"
	"testing"
)

func TestGetCollectionName(t *testing.T) {
	t.Run("Get struct Name", func(t *testing.T) {
		expected := "foo"
		found := getCollectionName(testhelper.Foo{})
		if expected != found {
			t.Errorf("Failed expected %v found %v", expected, found)
		}
	})
	t.Run("Get *struct Name", func(t *testing.T) {
		expected := "foo"
		found := getCollectionName(&testhelper.Foo{})
		if expected != found {
			t.Errorf("Failed expected %v found %v", expected, found)
		}
	})
	t.Run("Get []struct Name", func(t *testing.T) {
		expected := "foo"
		found := getCollectionName([]testhelper.Foo{})
		if expected != found {
			t.Errorf("Failed expected %v found %v", expected, found)
		}
	})
	t.Run("Get *[]struct Name", func(t *testing.T) {
		expected := "foo"
		found := getCollectionName(&[]testhelper.Foo{})
		if expected != found {
			t.Errorf("Failed expected %v found %v", expected, found)
		}
	})
	t.Run("Get String", func(t *testing.T) {
		expected := "foo"
		found := getCollectionName(expected)
		if expected != found {
			t.Errorf("Failed expected %v found %v", expected, found)
		}
	})

}
