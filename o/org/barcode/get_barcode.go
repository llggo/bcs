package barcode

func GetByID(id string) (*BarCode, error) {
	var b BarCode
	return &b, TableBarcode.ReadByID(id, &b)
}

func getBarCode(where map[string]interface{}) (*BarCode, error) {
	var b BarCode
	return &b, TableBarcode.ReadOne(where, &b)
}

func GetByCode(code string) (*BarCode, error) {
	var b BarCode
	return &b, TableBarcode.ReadOne(map[string]interface{}{
		"code":  code,
		"dtime": 0,
	}, &b)
}

func GetAll() ([]*BarCode, error) {
	var bs = []*BarCode{}
	return bs, TableBarcode.UnsafeReadAll(&bs)
}

func Count(where map[string]interface{}) (int, error) {
	return TableBarcode.UnsafeCount(where)
}
