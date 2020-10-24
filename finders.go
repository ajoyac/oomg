package oomg

func (d *Db) All(models interface{}) error {
	return Q(d).All(models)
}
func (q *Query) All(models interface{}) (err error) {

	opts := q.getFindOpts()
	filter := q.getFilter()
	collection := q.database.Collection(models)
	ctx, cancel := newCtx()
	defer cancel()

	cursor, err := collection.Find(ctx, filter, opts)
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
func (d *Db) One(model interface{}) error {
	return Q(d).One(model)
}
func (q *Query) One(model interface{}) (err error) {

	opts := q.getFindOneOpts()
	filter := q.getFilter()
	collection := q.database.Collection(model)
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
