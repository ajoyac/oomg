package oomg

//place holder for the creation options filler
type Creator struct {
	database *Db
}

func C(d *Db) *Creator {
	return &Creator{
		database: d,
	}
}
