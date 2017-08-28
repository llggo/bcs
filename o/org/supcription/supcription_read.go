package supcription

func GetByID(id string) (*Supcription, error) {
	var g Supcription
	return &g, SupcriptionTable.ReadByID(id, &g)
}

func getQrCode(where map[string]interface{}) (*Supcription, error) {
	var g Supcription
	return &g, SupcriptionTable.ReadOne(where, &g)
}

func GetByQrName(qr_name string) (*Supcription, error) {
	var g Supcription
	return &g, SupcriptionTable.ReadOne(map[string]interface{}{
		"qr_name": qr_name,
		"dtime":   0,
	}, &g)
}

func GetAll(where map[string]interface{}) ([]*Supcription, error) {
	var gs = []*Supcription{}
	return gs, SupcriptionTable.UnsafeReadMany(where, &gs)
}
