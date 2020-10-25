package oomg

func All(models interface{}) error {
	return Q(models).All(models)
}
func (q *Query) All(models interface{}) (err error) {

	opts := q.getFindOpts()
	filter := q.getFilter()

	if q.collection == nil {
		q.collection = collection(models)
	}

	ctx, cancel := newCtx()
	defer cancel()

	cursor, err := q.collection.Find(ctx, filter, opts)
	if err != nil {
		log.Error("Failed to Find model")
		return
	}
	err = cursor.All(ctx, models)
	if err != nil {
		log.Error("Failed to map query results and model")
		return
	}
	return
}
func One(model interface{}) error {
	return Q(model).One(model)
}
func (q *Query) One(model interface{}) (err error) {

	opts := q.getFindOneOpts()
	filter := q.getFilter()
	collection := q.collection

	ctx, cancel := newCtx()
	defer cancel()

	cursor := collection.FindOne(ctx, filter, opts)

	err = cursor.Decode(model)
	if err != nil {
		log.Error("Failed during Decode")
		return
	}
	return
}
