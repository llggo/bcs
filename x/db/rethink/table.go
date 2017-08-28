package rethink

import (
	r "gopkg.in/dancannon/gorethink.v2"
	"qrcode/pba/x/math"
	"qrcode/pba/x/web"
	"time"
)

type IdMaker interface {
	Next() string
}

type Table struct {
	IdMaker
	*UnsafeTable
}

var defaultIdMaker = &math.RandStringMaker{Prefix: "def", Length: 20}

func NewTable(db *Database, name string, IdMaker IdMaker) *Table {
	var t = &Table{
		UnsafeTable: NewUnsafeTable(db, name),
		IdMaker:     IdMaker,
	}
	if t.IdMaker == nil {
		t.IdMaker = defaultIdMaker
	}
	return t
}

func (t *Table) Create(i IModel) error {
	i.BeforeCreate()
	i.SetID(t.IdMaker.Next())
	return t.UnsafeInsert(i)
}

func (t *Table) UpdateByID(id string, i IModel) error {
	i.BeforeUpdate()
	i.SetID(id)
	return t.UnsafeUpdateByID(id, i)
}

func (t *Table) MarkDelete(id string) error {
	var data = map[string]interface{}{
		"dtime": time.Now().Unix(),
	}
	return t.UnsafeUpdateByID(id, data)
}

func (t *Table) ReadAll(ptr interface{}) error {
	return t.UnsafeReadMany(func(doc r.Term) r.Term {
		return doc.Field("dtime").Eq(0)
	}, ptr)
}

func (t *Table) ReadMany(key string, values []string, ptr interface{}) error {
	return t.UnsafeReadMany(func(doc r.Term) r.Term {
		return r.Expr(values).Contains(doc.Field(key)).And(doc.Field("dtime").Eq(0))
	}, ptr)
}

func (t *Table) ReadOne(where interface{}, ptr interface{}) error {
	return t.UnsafeReadOne(where, ptr)
}

func (t *Table) ReadByID(id string, ptr interface{}) error {
	return t.UnsafeReadOne(map[string]interface{}{
		"id":    id,
		"dtime": 0,
	}, ptr)
}

func (t *Table) NotExist(where map[string]interface{}) error {
	var c, err = t.UnsafeTable.UnsafeCount(where)
	if err != nil {
		return err
	}
	if c > 0 {
		return web.BadRequest("already exist")
	}
	return nil
}

func (t *Table) ReadByArrID(ids []string, ptr interface{}) error {
	return t.UnsafeRunGetAll(func(table r.Term) r.Term {
		return table.GetAll(r.Args(ids))
	}, ptr)
}
