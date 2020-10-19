package oomg

import "go.mongodb.org/mongo-driver/mongo"

type M map[string]interface{}

type Query struct {
	database *mongo.Database
	filter   M
}

func Q(d *Db) *Query {
	return &Query{
		database: d.db,
		filter:   M{},
	}
}
