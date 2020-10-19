package oomg

import log "github.com/sirupsen/logrus"

func (c *Db) All(models interface{}) error {
	return Q(c).All(models)
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
