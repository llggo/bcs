package rethink

import (
	"github.com/golang/glog"
	r "gopkg.in/dancannon/gorethink.v2"
)

var databases = map[string]*Database{}

//Database is rethink db
type Database struct {
	Session *r.Session
	Tables  []string
}

type FuncTerm func(r.Term) r.Term

func GetDB(name string) *Database {
	return databases[name]
}

//NewDB i
func NewDB(server, database string) {
	session, err := r.Connect(r.ConnectOpts{
		Address:  server,
		Database: database,
	})

	if err != nil {
		glog.Fatal("Connect to DB", err)
	}

	var tables []string
	cursor, err := r.TableList().Run(session)
	if err != nil {
		glog.Fatal("Table List Err", err)
	}
	defer cursor.Close()
	if err := cursor.All(&tables); err != nil {
		glog.Fatal("Read all tables", err)
	}

	databases[database] = &Database{
		Session: session,
		Tables:  tables,
	}
}

func CreateDBIfNotExist(address string, name string) error {
	var session, err = r.Connect(r.ConnectOpts{
		Address: address,
	})
	if err != nil {
		return err
	}
	var names []string
	cursor, err := r.DBList().Run(session)
	if err != nil {
		glog.Fatal("db list failed", err)
		return err
	}
	defer cursor.Close()
	if err := cursor.All(&names); err != nil {
		glog.Fatal("Create database if not exist", err)
		return err
	}

	for _, n := range names {
		if name == n {
			return nil
		}
	}

	_, err = r.DBCreate(name).Run(session)
	return err
}

func (db *Database) CreateTableIfNotExists(name string) error {
	for _, table := range db.Tables {
		if table == name {
			return nil
		}
	}
	return r.TableCreate(name).Exec(db.Session)
}

func (db *Database) Insert(table string, i interface{}) error {
	return r.Table(table).Insert(i).Exec(db.Session)
}

func (db *Database) UpdateByID(table string, id string, data interface{}) error {
	return r.Table(table).Get(id).Update(data).Exec(db.Session)
}

func (db *Database) DeleteByID(table string, id string) error {
	return r.Table(table).Get(id).Delete().Exec(db.Session)
}

func (db *Database) UpdateByIndex(table, index string, keys []interface{}, data interface{}) error {
	return r.Table(table).GetAllByIndex(index, keys).Update(data).Exec(db.Session)
}

func (db *Database) UpdateWhere(table string, where, data interface{}) error {
	return r.Table(table).Filter(where).Update(data).Exec(db.Session)
}

func (db *Database) Count(table string, where interface{}) (int, error) {
	cursor, err := r.Table(table).Filter(where).Count().Run(db.Session)
	if err != nil {
		return 0, err
	}
	defer cursor.Close()
	var count int
	cursor.One(&count)
	return count, nil
}

func (db *Database) IndexList(table string) ([]string, error) {
	var names = []string{}
	cursor, err := r.Table(table).IndexList().Run(db.Session)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	err = cursor.All(&names)
	return names, err
}

func (db *Database) IndexCreate(table, key string, index interface{}) error {

	var indexes, err = db.IndexList(table)
	if err != nil {
		return err
	}

	for _, v := range indexes {
		if v == key {
			return nil
		}
	}
	return r.Table(table).IndexCreateFunc(key, index).Exec(db.Session)
}
