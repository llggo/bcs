package qrcode

func GetByID(id string) (*QrCode, error) {
	var qr QrCode
	return &qr, TableQrcode.ReadByID(id, &qr)
}

func getQrCode(where map[string]interface{}) (*QrCode, error) {
	var qr QrCode
	return &qr, TableQrcode.ReadOne(where, &qr)
}

func GetByQrName(qr_name string) (*QrCode, error) {
	var qr QrCode
	return &qr, TableQrcode.ReadOne(map[string]interface{}{
		"qr_name": qr_name,
		"dtime":   0,
	}, &qr)
}

func GetAll(where map[string]interface{}) ([]*QrCode, error) {
	var qrs = []*QrCode{}
	return qrs, TableQrcode.UnsafeReadMany(where, &qrs)
}

func Count(where map[string]interface{}) (int, error) {
	return TableQrcode.UnsafeCount(where)
}
