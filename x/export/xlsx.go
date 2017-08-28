package export

import (
	"fmt"
	"os"
	"path/filepath"

	ex "github.com/tealeg/xlsx"
)

type Row struct {
	*ex.Row
}

type Sheet struct {
	*ex.Sheet
}

type Xlsx struct {
	FileName string
	*ex.File
}

func NewWriteXlsx(fileName string) *Xlsx {
	dir, _ := filepath.Split(fileName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return &Xlsx{
		FileName: fileName,
		File:     ex.NewFile(),
	}
}

func NewOpenFileXlsx(fileName string) (*Xlsx, error) {
	xlFile, err := ex.OpenFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Xlsx{
		FileName: fileName,
		File:     xlFile,
	}, nil
}

func (x *Xlsx) Save() {
	err := x.File.Save(x.FileName + ".xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func (x *Xlsx) NewSheet(name string) *Sheet {
	sheet, err := x.File.AddSheet(name)
	if err != nil {
		fmt.Println(err)
	}
	return &Sheet{
		Sheet: sheet,
	}
}

func (s *Sheet) NewRow() *Row {
	return &Row{
		Row: s.AddRow(),
	}
}

func (r *Row) AddCell(value ...string) {
	for _, v := range value {
		c := r.Row.AddCell()
		c.Value = v
		s := c.GetStyle()
		s.Alignment.WrapText = true
		s.Alignment.Horizontal = "center"
		s.Alignment.Vertical = "center"
		s.Font.Size = 11
	}

}
