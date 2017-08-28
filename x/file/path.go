package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Directory struct {
	Dir string
}

func NewDirectory(dir string) *Directory {
	return &Directory{
		Dir: dir,
	}
}

func (d *Directory) ListFile() []os.FileInfo {
	fi, err := ioutil.ReadDir(d.Dir)
	if err != nil {
		fmt.Println("libs/file - path: ", err)
		return nil
	}
	return fi
}
