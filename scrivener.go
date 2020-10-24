package oomg

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *Db) Create(models interface{}) error {
	return C(d).Create(models)
}
func (c *Creator) Create(model interface{}) (err error) {

	collection := c.database.Collection(model)
	ctx, cancel := newCtx()
	log.WithField("Collection", collection.Name()).Debug("Creating Document")
	defer cancel()

	insertResult, err := collection.InsertOne(ctx, model)

	if err != nil {
		log.Error("Failed to InsertOne")
		return
	}
	//This options its not perfomance wise
	// but it easy way to get the id into the model without direct modifying the
	// model if the user not create an ID field as required.
	//also StackOverFlow was down at the time.
	id, _ := insertResult.InsertedID.(primitive.ObjectID)

	log.Debug("Insert Result", insertResult)

	err = c.database.Filter(M{"_id": id}).One(model)
	if err != nil {
		log.Error("Failed to Retrive object created")
		log.Error(err)
		return
	}

	return
}
