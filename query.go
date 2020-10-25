package oomg

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Query struct {
	collection *mongo.Collection
	filter     M
}

func Q(m interface{}) *Query {
	return &Query{
		collection: collection(m),
		filter:     M{},
	}
}

func (q *Query) getFindOneOpts() *options.FindOneOptions {
	return options.FindOne()
}
func (q *Query) getFindOpts() *options.FindOptions {
	return options.Find()
}

func (q *Query) getFilter() bson.M {
	return bson.M(q.filter)
}

func Filter(m M) *Query {
	return Q(nil).Filter(m)
}
func (q *Query) Filter(m M) *Query {
	q.filter = m
	return q
}
func (q *Query) Collection(i interface{}) *Query {
	q.collection = collection(i)
	return q
}
