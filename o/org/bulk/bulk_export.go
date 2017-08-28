package bulk

import (
	"path/filepath"
	"qrcode-bulk/qrcode-bulk-generator/o/org/qrcode"
	"qrcode-bulk/qrcode-bulk-generator/o/org/verify_code"
	"qrcode-bulk/qrcode-bulk-generator/x/export"
)

func ExportExcel(bulkID string) (*export.Xlsx, error) {

	b, err := GetByID(bulkID)
	if err != nil {
		return nil, err
	}

	//init excel
	var filename = filepath.Join(dir, "2", b.Name+"_"+b.ID)
	var xlsx = export.NewWriteXlsx(filename)
	var sheet = xlsx.NewSheet("Code")
	var row = sheet.NewRow()
	row.AddCell("Verify Code", "Bulk Name", "Bulk QR Code Url")

	qrcode, err := qrcode.GetAll(map[string]interface{}{
		"bulk_id": b.ID,
	})
	if err != nil {
		return nil, err
	}

	for _, q := range qrcode {
		c, err := verify_code.GetAll(map[string]interface{}{
			"qrcode_id": q.ID,
		})

		if err != nil {
			return nil, err
		}

		for _, v := range c {
			row = sheet.NewRow()
			row.AddCell(v.VCode, b.Name, "http://mirascan.vn:31001/api/handle/welcome?qrcode_id="+q.GetID())
			// row.AddCell(v.VCode, b.Name, "http://"+config.Station().Server.Name+"/api/handle/wellcome?id="+q.ID)
		}
	}

	return xlsx, nil
}
