package sql

import (
	"database/sql/driver"
	"errors"
)

type NullString string

// sql serial
func (j NullString) Value() (driver.Value, error) {
	return string(j), nil
}

// sql scan
func (j *NullString) Scan(src interface{}) error {
	if src == nil {
		*j = ""
		return nil
	}
	var data, ok = src.([]byte)
	if ok {
		*j = NullString(string(data))
		return nil
	}
	return errors.New("scan non string into string")
}
