package oomg

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create a model into the DB
//if you are looking to save a none struct
//create a Scribe  calling with the S() function the name of collection
func Create(model interface{}) error {
	return S(model).Create(model)
}

func (s *Scribe) Create(model interface{}) (err error) {

	if s.collection == nil {
		s.collection = collection(model)
	}

	collection := s.collection
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

	err = Q(s.collection.Name()).Filter(M{"_id": id}).One(model)
	if err != nil {
		log.Error("Failed to Retrive object created")
		log.Error(err)
		return
	}

	return
}
