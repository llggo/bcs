package bulk

import (
	"fmt"
	"qrcode-bulk/qrcode-bulk-generator/o/org/qrcode"
	"qrcode-bulk/qrcode-bulk-generator/o/org/verify_code"
	"qrcode-bulk/qrcode-bulk-generator/x/math"
	"qrcode-bulk/qrcode-bulk-generator/x/web"
	"qrcode-bulk/qrcode-bulk-generator/x/work"
	"strings"
	"time"
)

// ChanCBQGV create bulk qrcode gen verify code
var BulkChan = make(chan (BulkType), 1000)

// ChanCBGQGVFile create bulk gen qrcode gen verify code
// var ChanCBGQGVFile = make(chan (BDGQGVFile), 1000)

// BDQGV struct bulk gen qrcode gen verify code
type BulkType struct {
	Type   int           `bson:"type" json:"type"`
	Number int64         `bson:"number" json:"number"`
	Bulk   Bulk          `bson:"bulk" json:"bulk"`
	QrCode qrcode.QrCode `bson:"qrcode" json:"qrcode"`
}

// BDGQGV struct bulk gen qrcode gen verify code
// type BDGQGVFile struct {
// 	Bulk   Bulk          `bson:"bulk" json:"bulk"`
// 	QrCode qrcode.QrCode `bson:"qrcode" json:"qrcode"`
// 	File   File          `bson:"file" json:"file"`
// }

// type File struct {
// 	Path    string        `bson:"path" json:"path"`
// 	GenType typeGenerator `bson:"gen_type" json:"gen_type"`
// }

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

//CreateBDQGV is Create Bulk type 1, 2, 3
func CreateBDQGV(b BulkType) {
	start := time.Now() // get current time
	fmt.Println(fmt.Sprintf("Create bulk [%v] type [%v] start", b.Bulk.Name, b.Type))

	//Create bulk
	switch b.Type {
	case 2:
		b.Bulk.QrCodeNumber = 1
		b.Bulk.VerifyNumber = b.Number
		b.Bulk.Type = 2
	case 3:
		b.Bulk.QrCodeNumber = b.Number
		b.Bulk.VerifyNumber = b.Number
		b.Bulk.Type = 3
	default:
		b.Bulk.QrCodeNumber = 1
		b.Bulk.VerifyNumber = 0
		b.Bulk.Type = 1
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
		}

		// var filename = filepath.Join(dir, "2", b.Bulk.BName+"_"+b.Bulk.ID)
		// var xlsx = export.NewWriteXlsx(filename)

		// if err != nil {
		// 	panic(err)
		// }
		// var sheet = xlsx.NewSheet("Code")

		// var row = sheet.NewRow()
		// row.AddCell("Verify Code", "Bulk Name", "Bulk QR Code Url")
		// row = sheet.NewRow()
		// row.AddCell("", b.Bulk.BName, "http://"+config.Station().Server.Name+"/api/handle/view?id="+b.QrCode.ID)

		s.Work = func(i int64) {

			//Create qrcode
			if b.Type == 3 {
				b.QrCode.BulkID = b.Bulk.ID
				b.QrCode.UserID = b.Bulk.UserID
				fmt.Println(b.QrCode)
				err = b.QrCode.Create()

				if err != nil {
					fmt.Println(err)
				}
			}

			//Create Verify Code
			var vc verify_code.VerifyCode
			vc.QrcodeID = b.QrCode.ID
			vc.VCode = strings.Trim(math.RandString("", 12), "_")
			// row = sheet.NewRow()
			// row.AddCell(vc.VCode)

			web.AssertNil(vc.Create())
		}

		s.Done = func() {
			// xlsx.Save()
			var log = fmt.Sprintf("Gen [%v] verify code of bulk [%v] type [%v] success", b.Bulk.VerifyNumber, b.Bulk.Name, 2)
			fmt.Println(log)
		}

		s.Run()
	}

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Create bulk [%v] type 2 done Run %s", b.Bulk.Name, elapsed))
}

// type typeGenerator string

// const GeneratorDefault = typeGenerator("default")

// const GeneratorAuto = typeGenerator("auto")

// const GeneratorOption = typeGenerator("option")

// //CreateBDGQGV is Create Bulk type 4
// func CreateBDGQGV(d BDGQGVFile) {
// 	start := time.Now() // get current time
// 	fmt.Println(fmt.Sprintf("Create bulk [%v] type 3 start", d.Bulk.BName))

// 	//Read Excel
// 	xlFile, err := export.NewOpenFileXlsx(d.File.Path)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	//Write Excel
// 	var filename = filepath.Join(dir, "3")
// 	var xlsx = export.NewWriteXlsx(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var wSheet = xlsx.NewSheet("Code")
// 	var row = wSheet.NewRow()
// 	row.AddCell(d.Bulk.BName)

// 	row = wSheet.NewRow()
// 	row.AddCell("QR Code ID", "Verify Code", "QR Code Url")

// 	//Create bulk
// 	var number = 0
// 	d.Bulk.VerifyNumber = 0
// 	d.Bulk.QrCodeNumber = 0
// 	err = d.Bulk.Create()

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// Create qrcode
// 	d.QrCode.BulkID = d.Bulk.ID
// 	d.QrCode.UserID = d.Bulk.UserID

// 	//Read Sheet
// 	for _, sheet := range xlFile.File.Sheets {

// 		var s = &work.Work{
// 			Number:    len(sheet.Rows),
// 			TaskLimit: 4000,
// 		}

// 		var t = -1
// 		var v = -1
// 		for i, c := range sheet.Rows[0].Cells {
// 			if strings.EqualFold(c.String(), "value") {
// 				v = i
// 			}
// 			if strings.EqualFold(c.String(), "type") {
// 				t = i
// 			}
// 		}

// 		//Read row by task WORK (Init)
// 		s.Work = func(i int) {

// 			//Read Cell
// 			var r = sheet.Rows[i]
// 			for index, c := range r.Cells {
// 				value := c.String()
// 				if i != 0 {
// 					switch d.File.GenType {
// 					//Gen qrcode (type by cell) & verify code auto filter by type & value column in sheet of excel
// 					case GeneratorAuto:
// 						d.QrCode.QrName = fmt.Sprintf("[%v] %v", i, d.QrCode.QrName)
// 						d.QrCode.QrTemplate = "default"
// 						if index == t {
// 							//Type
// 							d.QrCode.QrType = value
// 						}
// 						if index == v {
// 							// Value
// 							//Data qrcode
// 							d.QrCode.QrData = setQrData(d.QrCode.QrType, value)
// 						}

// 					case GeneratorOption:

// 					//Gen qrcode (type static) & verify code auto filter by value column in sheet of excel
// 					default:
// 						if index == v {
// 							// Value
// 							//Data qrcode
// 							fmt.Println(fmt.Sprintf("%v - %v \n", index, value))
// 							d.QrCode.QrData = setQrData(d.QrCode.QrType, value)
// 						}

// 					}
// 				}
// 			}

// 			//Create Qrcode
// 			err = d.QrCode.Create()
// 			if err != nil {
// 				fmt.Println(err)
// 			}

// 			//Create Verify Code
// 			var vc verify_code.VerifyCode
// 			vc.QrcodeID = d.QrCode.ID
// 			vc.VCode = strings.Trim(math.RandStringNumber("", 12), "_")

// 			web.AssertNil(vc.Create())

// 			row = wSheet.NewRow()
// 			row.AddCell(d.QrCode.ID, vc.VCode, "http://"+config.Station().Server.Name+"/api/handle/view?id="+d.QrCode.ID)

// 			//Count qrcode & verify code
// 			number++
// 		}

// 		//Read row finish
// 		s.Done = func() {
// 			d.Bulk.VerifyNumber = number
// 			d.Bulk.QrCodeNumber = number
// 			err = d.Bulk.Update(&d.Bulk)
// 			if err != nil {
// 				fmt.Println(err)
// 			}

// 			xlsx.FileName = filepath.Join(xlsx.FileName, fmt.Sprintf("[%v] %v - %v", d.File.GenType, d.Bulk.BName, d.Bulk.ID))
// 			xlsx.Save()

// 			var log = fmt.Sprintf("Gen [%v] verify code of bulk [%v] type [%v] success", d.Bulk.VerifyNumber, d.Bulk.BName, 3)
// 			fmt.Println(log)
// 		}

// 		//Run task
// 		s.Run()

// 	}

// 	elapsed := time.Since(start)
// 	fmt.Println(fmt.Sprintf("Create bulk [%v] type 3 done Run %s", d.Bulk.BName, elapsed))
// }

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
