package oomg


func (c *Db) All(models interface{}) error {
	return Q(c).All(models)
}
func (q *Query) All(models interface{}) error {
    return nil
}
