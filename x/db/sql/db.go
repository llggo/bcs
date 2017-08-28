package sql

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"qrcode/pba/x/web"
	"time"
)

type Database struct {
	dsn string
	DB  *sql.DB
}

func NewMySQLDB(host string, name string, account string) *Database {
	var dsn = fmt.Sprintf("%v@tcp(%v)/%v", account, host, name)
	var db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	glog.Infof("MySQL connected %v", dsn)

	return &Database{
		dsn: dsn,
		DB:  db,
	}
}

func (db *Database) ForEach(selector sq.SelectBuilder, factory RowFactory, iterator func(RowScannable)) error {

	var rows, err = selector.RunWith(db.DB).Query()

	if err != nil {
		var q, _, _ = selector.ToSql()
		glog.ErrorDepth(2, q, err)
		return web.InternalServerError("query foreach")
	}

	defer rows.Close()
	for rows.Next() {
		var r = factory()
		if err = rows.Scan(r.ScanList()...); err != nil {
			var q, _, _ = selector.ToSql()
			glog.ErrorDepth(2, q, err)
			return web.InternalServerError("scan failed")
		}
		iterator(r)
	}
	return nil
}

func (db *Database) Update(selector sq.UpdateBuilder, old RowUpdatable, new RowUpdatable) error {
	var update, err = old.UpdateMap(new)
	if err != nil {
		return err
	}
	selector = selector.SetMap(update)
	_, err = selector.RunWith(db.DB).Exec()
	if err != nil {
		var q, _, _ = selector.ToSql()
		glog.ErrorDepth(2, q, err)
		return web.InternalServerError("update failed")
	}
	return nil
}

func (db *Database) Insert(selector sq.InsertBuilder, row RowWritable) error {
	var write, err = row.WriteList()
	if err != nil {
		return err
	}
	selector = selector.Values(write...)
	_, err = selector.RunWith(db.DB).Exec()
	if err != nil {
		var q, _, _ = selector.ToSql()
		glog.ErrorDepth(2, q, err)
		return web.InternalServerError("insert failed")
	}
	return nil
}

func (db *Database) MarkDelete(table string, id string) error {
	var where = map[string]interface{}{"id": id, "dtime": 0}
	var now = time.Now().Unix()
	var updateQuery = sq.Update(table).Where(where).SetMap(map[string]interface{}{
		"mtime": now,
		"dtime": now,
	})
	_, err := updateQuery.RunWith(db.DB).Exec()
	if err != nil {
		var q, _, _ = updateQuery.ToSql()
		glog.ErrorDepth(2, q, err)
		return web.InternalServerError("update failed")
	}
	return nil
}

func (d *Database) Count(table string, where sq.And) (int, error) {
	var countQuery = sq.Select("COUNT(id)").From(table).Where(where).Limit(1)
	var count int
	if err := countQuery.RunWith(d.DB).Scan(&count); err != nil {
		var q, _, _ = countQuery.ToSql()
		glog.ErrorDepth(2, q, err)
		return 0, web.InternalServerError("check new branch code")
	}
	return count, nil
}

func (d *Database) ReadOne(selector sq.SelectBuilder, factory RowFactory) (RowScannable, error) {
	var res RowScannable
	var err = d.ForEach(selector, factory, func(row RowScannable) {
		res = row
	})
	if err == nil && res == nil {
		return nil, web.BadRequest("data not found")
	}
	return res, err
}
