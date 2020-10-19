package oomg

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type M map[string]interface{}

type Query struct {
	database *Db
	filter   M
}

func Q(d *Db) *Query {
	return &Query{
		database: d,
		filter:   M{},
	}
}

func (q *Query) getFindOpts() *options.FindOptions {
	return options.Find()
}

func (q *Query) getFilter() bson.M {
	return bson.M(q.filter)
}
