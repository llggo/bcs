package model

import (
	"os"
	"qrcode-bulk/qrcode-bulk-generator/config/cons"
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
	"qrcode-bulk/qrcode-bulk-generator/x/math"
)

type TableWithCode struct {
	*mgo.Table
}

type TableWithTag struct {
	*mgo.Table
}

func NewTable(name string, idPrefix string) *mgo.Table {
	var db = GetDB()
	var idMaker = math.RandStringMaker{Prefix: idPrefix, Length: 20}
	return mgo.NewTable(db, name, &idMaker)
}

func NewTableWithCode(name string, idPrefix string) *TableWithCode {
	var table = NewTable(name, idPrefix)
	return &TableWithCode{Table: table}
}

func NewTableWithTag(name string, idPrefix string) *TableWithTag {
	var table = NewTable(name, idPrefix)
	return &TableWithTag{Table: table}
}

func GetDB() *mgo.Database {
	return mgo.GetDB(os.Getenv(cons.ENV_OBJECT_DB))
}
