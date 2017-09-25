package bulk

import (
	"fmt"
	"path/filepath"
	"bar-code/bcs/o/org/qrcode"
	"bar-code/bcs/o/org/verify_code"
	"bar-code/bcs/x/export"
	"bar-code/bcs/x/work"
)

func ExportExcel(bulkID string) (*export.Xlsx, error) {

	b, err := GetByID(bulkID)
	if err != nil {
		return nil, err
	}

	//init excel
	var filename = filepath.Join(dir, "3", b.Name+"_"+b.ID)
	var xlsx = export.NewWriteXlsx(filename)
	var sheet = xlsx.NewSheet("Code")
	var row = sheet.NewRow()
	if b.Type == 3 {
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
	} else if b.Type == 4 {
		row.AddCell("Bulk Name", "Bulk QR Code url")
		qrcode, err := qrcode.GetAll(map[string]interface{}{
			"bulk_id": b.GetID(),
		})
		if err != nil {
			return nil, err
		}
		for _, q := range qrcode {
			row = sheet.NewRow()
			row.AddCell(b.GetName(), "http://mirascan.vn:31001/api/handle/welcome?qrcode_id="+q.GetID())
		}
	}

	return xlsx, nil
}

func UpdateVerify(bulkID string) error {
	b, err := GetByID(bulkID)
	if err != nil {
		fmt.Println(err)
	}
	if b.Type == 4 {
		qrs, err := qrcode.GetAll(map[string]interface{}{
			"bulk_id": b.GetID(),
		})
		if err != nil {
			return err
		}
		var s = &work.Work{
			Number:    b.QrCodeNumber,
			TaskLimit: 4000,
			Async:     true,
		}
		s.Work = func(i int64) {
			qrs[i].VerifyLimit = b.VerifyLimit
			qrs[i].Update(qrs[i])
		}
		s.Done = func() {
			var log = fmt.Sprintln("Update verify of QR-Code Done!!!!")
			fmt.Println(log)
		}
		s.Run()
	}
	return nil
}
