package oomg

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Query struct {
	database   *Db
	filter     M
	collection string
}

func Q(d *Db) *Query {
	return &Query{
		database: d,
		filter:   M{},
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

func (c *Db) Filter(m M) *Query {
	return Q(c).Filter(m)
}
func (q *Query) Filter(m M) *Query {
	q.filter = m
	return q
}
func (q *Query) Collection(i interface{}) *Query {
	q.collection = getCollectionName(i)
	return q
}

func (q *Query) getCollection(i interface{}) *mongo.Collection {

	var collection *mongo.Collection

	if q.collection != "" {
		collection = q.database.Collection(q.collection)

	} else {
		collection = q.database.Collection(i)
	}

	return collection

}
