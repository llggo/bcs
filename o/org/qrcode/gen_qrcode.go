package qrcode

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"qrcode-bulk/qrcode-bulk-generator/x/math"
	qrcode "qrcode/qrcodelib"
	"strconv"
	"time"
)

var size = 256

func (qr *QrCode) CreateImage(content string) {
	colorFG := "#000000"
	qrt, err := qrcode.New(content, qrcode.High, colorFG)
	if err != nil {
		fmt.Println(err)
	}
	//checkError(err)
	png, err := qrt.PNG(size)
	pathfile, pathimg := qr.CreateNewFolderImg() // duong dan tuong doi,
	nameimg := math.RandString("img_", 20) + ".png"
	pathimage := pathfile + nameimg
	qr.QrPathImg = pathimg + nameimg
	fh, err := os.Create(pathimage)
	qr.QrSize = len(qrt.Bitmap())
	checkError(err)
	fh.Write(png)
	defer fh.Close()
	/*******************create img base64**********************/
	imgFile, err := os.Open(pathimage)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer imgFile.Close()
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	// Embed into an html without PNG file
	// img2html := "<html><body><img src=\"data:image/png;base64," + imgBase64Str + "\" /></body></html>"
	qr.QrPathBase64 = "data:image/png;base64," + imgBase64Str
	// w.Write([]byte(fmt.Sprintf(img2html)))
	/*******************create img base64**********************/
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func (qr *QrCode) CreateNewFolderImg() (string, string) {
	var err error
	t := time.Now().Local()
	x := t.Month()
	month := int(x)
	sYear := strconv.Itoa(t.Year())
	sMonth := strconv.Itoa(month)
	sDay := strconv.Itoa(t.Day())
	sHour := strconv.Itoa(t.Hour())
	sMinute := strconv.Itoa(t.Minute())
	// s_second := strconv.Itoa(t.Second())

	folder := []string{"static", "img", "qrcode", sYear, sMonth, sDay, sHour, sMinute}
	pathFolder := ""
	pathImg := ""
	var countFolder = len(folder)
	for i := 0; i < countFolder; i++ {
		pathFolder += folder[i] + "/"
		pathImg += folder[i] + "/"
		os.Mkdir(pathFolder, os.ModeSticky|0755)
		if err != nil {
			log.Println("Error creating directory")
			log.Println(err)
		}
	}
	return pathFolder, pathImg
}
