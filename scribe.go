package oomg

import "go.mongodb.org/mongo-driver/mongo"

type Scribe struct {
	collection *mongo.Collection
}

func S(model interface{}) *Scribe {
	return &Scribe{
		collection: collection(model),
	}

}
