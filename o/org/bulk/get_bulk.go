package bulk

func GetByID(id string) (*Bulk, error) {
	var b Bulk
	return &b, TableBulk.ReadByID(id, &b)
}

func getBulk(where map[string]interface{}) (*Bulk, error) {
	var b Bulk
	return &b, TableBulk.ReadOne(where, &b)
}

func GetByBName(bname string) (*Bulk, error) {
	var b Bulk
	return &b, TableBulk.ReadOne(map[string]interface{}{
		"name":  bname,
		"dtime": 0,
	}, &b)
}

func GetAll(where map[string]interface{}) ([]*Bulk, error) {
	var bs = []*Bulk{}
	return bs, TableBulk.UnsafeReadMany(where, &bs)
}

func Count(where map[string]interface{}) (int, error) {
	return TableBulk.UnsafeCount(where)
}
