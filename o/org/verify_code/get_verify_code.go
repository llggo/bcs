package verify_code

func GetByID(id string) (*VerifyCode, error) {
	var b VerifyCode
	return &b, TableVerifyCode.ReadByID(id, &b)
}

func getBulk(where map[string]interface{}) (*VerifyCode, error) {
	var b VerifyCode
	return &b, TableVerifyCode.ReadOne(where, &b)
}

func GetByBName(bname string) (*VerifyCode, error) {
	var b VerifyCode
	return &b, TableVerifyCode.ReadOne(map[string]interface{}{
		"name":  bname,
		"dtime": 0,
	}, &b)
}

func Find(where map[string]interface{}) (*VerifyCode, error) {
	var b VerifyCode
	return &b, TableVerifyCode.ReadOne(where, &b)
}

func GetAll(where map[string]interface{}) ([]*VerifyCode, error) {
	var bs = []*VerifyCode{}
	return bs, TableVerifyCode.UnsafeReadMany(where, &bs)
}

func Count(where map[string]interface{}) (int, error) {
	return TableVerifyCode.UnsafeCount(where)
}
