package customize

func GetByID(id string) (*Customize, error) {
	var cus Customize
	return &cus, TableCustomize.ReadByID(id, &cus)
}

func GetByQrID(qrid string) (*Customize, error) {
	var cus Customize
	return &cus, TableCustomize.ReadOne(map[string]interface{}{"qrcode_id": qrid}, &cus)
}

func getCusCode(where map[string]interface{}) (*Customize, error) {
	var cus Customize
	return &cus, TableCustomize.ReadOne(where, &cus)
}

func GetAll(where map[string]interface{}) ([]*Customize, error) {
	var cus = []*Customize{}
	return cus, TableCustomize.UnsafeReadMany(where, &cus)
}

func Count(where map[string]interface{}) (int, error) {
	return TableCustomize.UnsafeCount(where)
}
