package xlsx

import (
	"os"
	"path/filepath"

	"github.com/xuri/excelize"
)

var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

type col struct {
	Index int
	Name  string
}

func (c *col) Next() {
	var l = len(alphabet)

	if c.Index > l {

	}

	c.Name = alphabet[c.Index]
}

type Xlsx struct {
	Name string
	*excelize.File
}

type Sheet struct {
	Index int
	Name  string
	*excelize.File
}

type Row struct {
}

func NewSlsx(fileName string) *Xlsx {
	dir, _ := filepath.Split(fileName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return &Xlsx{
		Name: fileName,
		File: excelize.NewFile(),
	}
}

func (x *Xlsx) NewSheet(index int, name string) *Sheet {
	x.File.NewSheet(index, name)
	return &Sheet{
		Index: index,
		Name:  name,
		File:  x.File,
	}
}
