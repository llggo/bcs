package file

import "io/ioutil"
import "path/filepath"

type file struct {
	Dir      string
	FileName string
}

func NewFile(dir string, filename string) *file {
	return &file{
		Dir:      dir,
		FileName: filename,
	}
}

func (f *file) Read() ([]byte, error) {
	var p = filepath.Join(f.Dir, f.FileName)
	d, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (f *file) Write(d []byte) error {
	var p = filepath.Join(f.Dir, f.FileName)
	err := ioutil.WriteFile(p, d, 0644)
	if err != nil {
		return err
	}
	return nil
}
