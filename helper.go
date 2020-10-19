package oomg

import (
	"reflect"
	"strings"
)

func getCollectionName(c interface{}) string {

	t := reflect.TypeOf(c)

	if t.Kind() == reflect.Ptr {

		t = t.Elem()

	}

	switch t.Kind() {

	case reflect.Slice, reflect.Array:
		el := t.Elem()
		return strings.ToLower(el.Name())
	case reflect.String:
		v, _ := c.(string)
		return v
	default:
		return strings.ToLower(t.Name())

	}
}
