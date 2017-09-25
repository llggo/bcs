package bulk

import (
	"fmt"
	"bar-code/bcs/o/org/qrcode"
	"bar-code/bcs/o/org/verify_code"
	"bar-code/bcs/x/math"
	"bar-code/bcs/x/web"
	"bar-code/bcs/x/work"
	"strings"
	"time"
)

// BulkChan create bulk qrcode gen verify code
var BulkChan = make(chan (BulkType), 1000)

// BulkType struct bulk gen qrcode gen verify code
type BulkType struct {
	Type   int           `bson:"type" json:"type"`
	Number int64         `bson:"number" json:"number"`
	Bulk   Bulk          `bson:"bulk" json:"bulk"`
	QrCode qrcode.QrCode `bson:"qrcode" json:"qrcode"`
}

func init() {
	go compute()
}

func compute() {
	for {
		select {
		case d := <-BulkChan:
			go CreateBDQGV(d)
			// case d := <-ChanCBGQGVFile:
			// 	go CreateBDGQGV(d)
		}
	}
}

var dir = "static/export/bulk"

//CreateBDQGV is Create Bulk type 1, 2, 3, 4

const (
	// Define Ma xac thuc thong tin
	MA_XAC_THUC_THONG_TIN = 1

	//define Ma xac thuc bien doi
	MA_XAC_THUC_BIEN_DOI = 2

	// Define qrcode bien doi va xac thuc
	MA_QRCODE_BIEN_DOI_XAC_THUC = 3

	//define qrcode bien doi
	MA_QRCODE_BIEN_DOI = 4
)

func CreateBDQGV(b BulkType) {
	start := time.Now() // get current time
	fmt.Println(fmt.Sprintf("Create bulk [%v] type [%v] start", b.Bulk.Name, b.Type))

	//Create bulk
	switch b.Type {
	case MA_XAC_THUC_THONG_TIN:
		b.Bulk.QrCodeNumber = 1
		b.Bulk.VerifyNumber = 0
		b.Bulk.Type = 1
	case MA_XAC_THUC_BIEN_DOI:
		b.Bulk.QrCodeNumber = 1
		b.Bulk.VerifyNumber = b.Number
		b.Bulk.Type = 2
	case MA_QRCODE_BIEN_DOI_XAC_THUC:
		b.Bulk.QrCodeNumber = b.Number
		b.Bulk.VerifyNumber = b.Number
		b.Bulk.Type = 3
	case MA_QRCODE_BIEN_DOI:
		b.Bulk.QrCodeNumber = b.Number
		b.Bulk.VerifyNumber = 0
		b.Bulk.Type = 4
	default:
		fmt.Println(fmt.Sprintf("Default create bulk qrcode"))
	}

	err := b.Bulk.Create()
	if err != nil {
		fmt.Println(err)
	}

	//Create qrcode
	if b.Type == 2 || b.Type == 1 {
		b.QrCode.BulkID = b.Bulk.ID
		b.QrCode.UserID = b.Bulk.UserID

		err = b.QrCode.Create()

		if err != nil {
			fmt.Println(err)
		}
	}

	//Event create Verify Code
	if b.Type != 1 {
		var s = &work.Work{
			Number:    b.Number,
			TaskLimit: 4000,
			Async:     false,
		}

		s.Work = func(i int64) {

			//Create qrcode
			if MA_QRCODE_BIEN_DOI_XAC_THUC == b.Type || MA_QRCODE_BIEN_DOI == b.Type {
				b.QrCode.BulkID = b.Bulk.ID
				b.QrCode.UserID = b.Bulk.UserID
				b.QrCode.VerifyLimit = b.Bulk.VerifyLimit
				err = b.QrCode.Create()

				// fmt.Println(fmt.Sprintf("%v - %v - %v", time.Now(), i, b.QrCode.ID))

				if err != nil {
					fmt.Println(err)
				}
			}

			//Create Verify Code
			if MA_XAC_THUC_BIEN_DOI == b.Type || MA_QRCODE_BIEN_DOI_XAC_THUC == b.Type {
				var vc verify_code.VerifyCode
				vc.QrcodeID = b.QrCode.ID
				vc.VCode = strings.Trim(math.RandString("", 12), "_")
				web.AssertNil(vc.Create())
			}
			// row = sheet.NewRow()
			// row.AddCell(vc.VCode)

		}

		s.Done = func() {
			// xlsx.Save()
			var log = fmt.Sprintf("Gen [%v] verify code of bulk [%v] type [%v] success", b.Bulk.VerifyNumber, b.Bulk.Name, b.Type)
			fmt.Println(log)
		}

		s.Run()
	}

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Create bulk [%v] type [%v] done Run %s", b.Bulk.Name, b.Type, elapsed))
}

func setQrData(t, d string) map[string]interface{} {
	switch t {
	case "url":
		return map[string]interface{}{
			"url": d,
		}
	default:
		return map[string]interface{}{
			"content": d,
		}
	}
}
