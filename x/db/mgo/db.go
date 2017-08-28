package mgo

import mgo "gopkg.in/mgo.v2"

//Database is mongo db
type Database struct {
	*mgo.Database
}

var databases = map[string]*Database{}

func GetDB(name string) *Database {
	return databases[name]
}

//NewDB i
func NewDB(server, database string) (*Database, error) {
	session, err := mgo.Dial(server)
	databases[database] = &Database{
		Database: session.DB(database),
	}
	return databases[database], err
}

func (db *Database) Insert(table string, i interface{}) error {
	return db.C(table).Insert(i)
}

func (db *Database) UpdateByID(table string, id string, data interface{}) error {
	return db.C(table).UpdateId(id, data)
}

func (db *Database) DeleteByID(table string, id string) error {
	return db.C(table).RemoveId(id)
}

func (db *Database) UpdateWhere(table string, where, data interface{}) error {
	return db.C(table).Update(where, data)
}

func (db *Database) Count(table string, where interface{}) (int, error) {
	return db.C(table).Find(where).Count()
}

func (db *Database) Indexes(table string) ([]mgo.Index, error) {
	return db.C(table).Indexes()
}

func (db *Database) EnsureIndex(table string, index mgo.Index) error {
	return db.C(table).EnsureIndex(index)
}
