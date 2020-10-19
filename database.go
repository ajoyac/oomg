package oomg

import "go.mongodb.org/mongo-driver/mongo"

type Db struct {
	database *mongo.Database
}

func DB() *Db {
	return &Db{database: currentDB}
}

//Return a empty Query
func (d *Db) Q() *Query {
	return Q(d)
}

func (d *Db) Collection(i interface{}) *mongo.Collection {
	name := getCollectionName(i)
	return d.database.Collection(name)
}
