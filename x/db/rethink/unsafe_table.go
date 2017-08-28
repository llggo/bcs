package rethink

import (
	"github.com/golang/glog"
	r "gopkg.in/dancannon/gorethink.v2"
	"qrcode/pba/x/web"
)

type UnsafeTable struct {
	DB      *Database
	Name    string
	Indices []string
}

func NewUnsafeTable(db *Database, Name string) *UnsafeTable {
	err := db.CreateTableIfNotExists(Name)
	if err != nil {
		panic(err)
	}
	return &UnsafeTable{
		DB:   db,
		Name: Name,
	}
}

func (t *UnsafeTable) RawTable() r.Term {
	return r.Table(t.Name)
}

func (t *UnsafeTable) UnsafeChanges() (*r.Cursor, error) {
	res, err := r.Table(t.Name).Changes().Run(t.DB.Session)
	if err != nil {
		glog.ErrorDepth(2, err)
		return nil, web.InternalServerError("track change failed")
	}
	return res, nil
}

const errRecordNotFound = web.NotFound("record not found")
const errReadDataFailed = web.InternalServerError("read data failed")
const errInsertDataFailed = web.InternalServerError("insert data failed")
const errUpdateDataFailed = web.InternalServerError("update data failed")
const errCountDataFailed = web.InternalServerError("count data failed")
const errNoOutput = web.InternalServerError("no ouput for data")

func (t *UnsafeTable) UnsafeRunGetAll(f func(table r.Term) r.Term, ptr interface{}) error {
	var cursor, err = f(r.Table(t.Name)).Run(t.DB.Session)
	if err != nil {
		glog.ErrorDepth(3, err)
		return errReadDataFailed
	}
	if ptr == nil {
		return errNoOutput
	}
	defer cursor.Close()
	err = cursor.All(ptr)
	if err != nil {
		glog.ErrorDepth(3, err)
		return errReadDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeRunGetOne(f func(table r.Term) r.Term, ptr interface{}) error {
	var cursor, err = f(r.Table(t.Name)).Run(t.DB.Session)
	if err != nil {
		glog.ErrorDepth(3, err)
		return errReadDataFailed
	}
	if ptr == nil {
		return errNoOutput
	}
	defer cursor.Close()
	err = cursor.One(ptr)
	if err != nil {
		if err == r.ErrEmptyResult {
			return errRecordNotFound
		}

		glog.ErrorDepth(3, err)
		return errReadDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeInsert(obj interface{}) error {
	err := t.DB.Insert(t.Name, obj)
	if err != nil {
		glog.ErrorDepth(2, err)
		return errInsertDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeCount(where interface{}) (int, error) {
	count, err := t.DB.Count(t.Name, where)
	if err != nil {
		glog.ErrorDepth(2, err)
		return 0, errCountDataFailed
	}
	return count, nil
}

func (t *UnsafeTable) UnsafeGetByID(id string, ptr interface{}) error {
	return t.UnsafeRunGetOne(func(table r.Term) r.Term {
		return table.Get(id)
	}, ptr)
}

func (t *UnsafeTable) UnsafeUpdateByID(id string, data interface{}) error {
	err := t.DB.UpdateByID(t.Name, id, data)
	if err != nil {
		glog.ErrorDepth(2, err)
		return errUpdateDataFailed
	}
	return nil
}
func (t *UnsafeTable) UnsafeDeleteByID(id string) error {
	return t.DB.DeleteByID(t.Name, id)
}

func (t *UnsafeTable) UnsafeUpdateByIndex(index string, keys []interface{}, data interface{}) error {
	err := t.DB.UpdateByIndex(t.Name, index, keys, data)
	if err != nil {
		glog.ErrorDepth(2, err)
		return errUpdateDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeUpdateWhere(where, data interface{}) error {
	err := t.DB.UpdateWhere(t.Name, where, data)
	if err != nil {
		glog.ErrorDepth(2, err)
		return errUpdateDataFailed
	}
	return nil
}

func (t *UnsafeTable) UnsafeReadAll(ptr interface{}) error {
	return t.UnsafeRunGetAll(func(table r.Term) r.Term {
		return table
	}, ptr)
}

func (t *UnsafeTable) UnsafeReadMany(where interface{}, ptr interface{}) error {
	return t.UnsafeRunGetAll(func(table r.Term) r.Term {
		return table.Filter(where)
	}, ptr)

}

func (t *UnsafeTable) UnsafeReadOne(where interface{}, ptr interface{}) error {
	return t.UnsafeRunGetOne(func(table r.Term) r.Term {
		return table.Filter(where)
	}, ptr)
}

func (t *UnsafeTable) IndexCreateArray(key string, field []string) error {
	var temp = []interface{}{}
	for _, f := range field {
		temp = append(temp, r.Row.Field(f))
	}
	return t.DB.IndexCreate(t.Name, key, temp)
}

func (t *UnsafeTable) EnsureIndex(field string) error {
	return t.DB.IndexCreate(t.Name, field, r.Row.Field(field))
}

func (t *UnsafeTable) IsErrNotFound(err error) bool {
	return err == errRecordNotFound
}
