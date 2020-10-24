package oomg

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
