package barcode

import (
	"fmt"
	"image/png"
	"net/http"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func (b *BarServer) CreateImage(w http.ResponseWriter, r *http.Request) {
	a, err := code128.Encode("http://tinhte.vn")
	fmt.Print("err:", err)
	// Scale the barcode to 200x200 pixels
	c, err := barcode.Scale(a, 300, 100)
	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, c)
}
