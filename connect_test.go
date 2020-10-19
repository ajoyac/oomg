package oomg

import "testing"

func TestConnect(t *testing.T) {
	t.Run("RealURL", func(t *testing.T) {
		err := Connect(ConnectOptions{
			URL: "mongodb://localhost:27017",
		})
		if err != nil {
			t.Error(err)
		}
	})

}
