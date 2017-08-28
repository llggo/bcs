package model

import (
	"qrcode-bulk/qrcode-bulk-generator/x/db/mgo"
)

type WithTag struct {
	mgo.BaseModel `bson:",inline"`
	Type          string `bson:"type" json:"type"`
	Tag           string `bson:"tag" json:"tag"`
}

type IWithTag interface {
	mgo.IModel
}

func (t *TableWithTag) GetByType(types []string, ptr interface{}) error {
	return t.ReadMany("type", types, ptr)
}
