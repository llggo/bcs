package rethink

import (
	"github.com/golang/glog"
	"reflect"
)

type CacheTable struct {
	*Table
	ModelType reflect.Type
	OnCreated func(IModel)
	OnUpdated func(IModel)
	OnDeleted func(IModel)
	data      map[string]IModel
}

func (c *CacheTable) makeArray() reflect.Value {
	return reflect.MakeSlice(reflect.SliceOf(c.ModelType), 0, 0)
}

func (c *CacheTable) makeMap() reflect.Value {
	return reflect.MakeMap(reflect.MapOf(reflect.TypeOf(""), c.ModelType))
}

func NewCacheTable(t *Table, ptr IModel) *CacheTable {
	ptrType := reflect.TypeOf(ptr)
	var f = func(IModel) {}
	var c = &CacheTable{
		Table:     t,
		ModelType: ptrType,
		OnCreated: f,
		OnUpdated: f,
		OnDeleted: f,
		data:      map[string]IModel{},
	}
	return c
}

func (c *CacheTable) Refresh() error {
	var arr = c.makeArray()
	x := reflect.New(arr.Type())
	x.Elem().Set(arr)
	err := c.ReadAll(x.Interface())
	if err != nil {
		return err
	}

	arr = x.Elem()

	for i := 0; i < arr.Len(); i++ {
		var d = arr.Index(i).Interface().(IModel)
		c.data[d.GetID()] = d
	}
	return nil
}

func (c *CacheTable) Watch() {
	watcher, err := c.UnsafeChanges()
	if err != nil {
		glog.Error("watch change", c.Name, c)
		return
	}

	change := c.makeMap()
	for watcher.Next(change) {
		old := change.FieldByName("old")
		new := change.FieldByName("new").Interface().(IModel)
		if old.IsNil() {
			c.data[new.GetID()] = new
			c.OnCreated(new)
		} else {
			if new.IsDeleted() {
				delete(c.data, new.GetID())
				c.OnDeleted(new)
			} else {
				c.data[new.GetID()] = new
				c.OnUpdated(new)
			}
		}
	}
}

func (c *CacheTable) GetByID(id string) IModel {
	return c.data[id]
}

func (c *CacheTable) Data() map[string]IModel {
	return c.data
}

func (c *CacheTable) DataArray() []IModel {
	var res = []IModel{}
	for _, d := range c.data {
		res = append(res, d)
	}
	return res
}
