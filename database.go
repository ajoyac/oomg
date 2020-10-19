package oomg

import "go.mongodb.org/mongo-driver/mongo"

type Db struct {
	db *mongo.Database
}

func DB() *Db {
	return &Db{db: currentDB}
}

//Return a empty Query
func (d *Db) Q() *Query {
	return Q(d)
}
