package sql

import (
	sq "github.com/Masterminds/squirrel"
)

type Table struct {
	DB      *Database
	Name    string
	Columns []string
	Factory RowFactory
}

func (db *Database) NewTable(name string, columns []string, factory RowFactory) *Table {
	return &Table{
		DB:      db,
		Name:    name,
		Columns: columns,
		Factory: factory,
	}
}

func (t *Table) ForEach(where map[string]interface{}, iterator func(RowScannable)) error {
	var selector = sq.Select(t.Columns...).From(t.Name).Where(where)
	return t.DB.ForEach(selector, t.Factory, iterator)
}

func (t *Table) MarkDelete(id string) error {
	return t.DB.MarkDelete(t.Name, id)
}

func (t *Table) Create(v RowWritable) error {
	var insertQuery = sq.Insert(t.Name).Columns(t.Columns...)
	return t.DB.Insert(insertQuery, v)
}

func (t *Table) Update(where map[string]interface{}, old RowUpdatable, new RowUpdatable) error {
	var updateQuery = sq.Update(t.Name).Where(where)
	return t.DB.Update(updateQuery, old, new)
}

func (t *Table) ReadOne(where map[string]interface{}) (RowScannable, error) {
	var selector = sq.Select(t.Columns...).From(t.Name).Where(where)
	return t.DB.ReadOne(selector, t.Factory)
}

func (t *Table) UpdateByID(id string, old RowUpdatable, new RowUpdatable) error {
	var updateQuery = sq.Update(t.Name).Where(map[string]interface{}{"id": id})
	return t.DB.Update(updateQuery, old, new)
}

func (t *Table) Count(where map[string]interface{}) (int, error) {
	return t.DB.Count(t.Name, sq.And{sq.Eq(where)})
}
