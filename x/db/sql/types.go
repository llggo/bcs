package sql

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type MapSQL map[string]interface{}

// sql serial
func (j MapSQL) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// sql scan
func (j *MapSQL) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	var data, ok = src.([]byte)
	if ok {
		return json.Unmarshal(data, j)
	}
	return errors.New("scan non []byte into json")
}

type ArrInt64SQL []int64

func (j ArrInt64SQL) MarshalJSON() ([]byte, error) {
	var str = make([]string, len(j))
	for i, v := range j {
		str[i] = fmt.Sprintf("%v", v)
	}
	return json.Marshal(str)
}

func (j *ArrInt64SQL) UnmarshalJSON(data []byte) error {
	var str = make([]json.Number, 0)
	var err = json.Unmarshal(data, &str)

	if err != nil {
		return err
	}
	var arr = make([]int64, len(str))
	for i, s := range str {
		if arr[i], err = s.Int64(); err != nil {
			return err
		}
	}

	*j = arr

	return nil
}

// sql serial
func (j ArrInt64SQL) Value() (driver.Value, error) {
	var data, err = json.Marshal([]int64(j))
	return data, err
}

// sql scan
func (j *ArrInt64SQL) Scan(src interface{}) error {
	var data, ok = src.([]byte)
	if ok {
		var arr = make([]int64, 0)
		if err := json.Unmarshal(data, &arr); err != nil {
			return err
		}
		*j = ArrInt64SQL(arr)
		return nil
	}
	return errors.New("scan non []byte into array int64")
}

type ArrStrSQL []string

// sql serial
func (j ArrStrSQL) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// sql scan
func (j *ArrStrSQL) Scan(src interface{}) error {
	var data, ok = src.([]byte)
	if ok {
		return json.Unmarshal(data, j)
	}
	return errors.New("scan non []byte into an array of string")
}
