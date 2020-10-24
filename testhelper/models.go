package testhelper

import "go.mongodb.org/mongo-driver/bson/primitive"

type Foo struct {
	ID primitive.ObjectID `bson:"omitempty"`
}
