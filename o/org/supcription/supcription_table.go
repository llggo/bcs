package supcription

import (
	"errors"
	"qrcode-bulk/qrcode-bulk-generator/o/model"

	"gopkg.in/mgo.v2/bson"
)

var SupcriptionTable = model.NewTable("supcription", "scp")

func NewSupcription() string {
	return SupcriptionTable.Next()
}

func (g *Supcription) Create() error {
	return SupcriptionTable.Create(g)
}

func MarkDelete(id string) error {
	return SupcriptionTable.MarkDelete(id)
}
func (g *Supcription) Update(newValue *Supcription) error {
	return SupcriptionTable.UnsafeUpdateByID(g.ID, newValue)
}

func AllQrCodeShare(branchID string) ([]Supcription, error) {
	var res = make([]Supcription, 0)
	query := SupcriptionTable.C().Find(bson.M{"branch_id": branchID})
	iter := query.Iter()
	if nil == iter {
		return nil, errors.New("\tERROR: Group Table.go::AllGroup")
	}
	var c Supcription
	for iter.Next(&c) {
		res = append(res, c)
	}
	return res, nil
}
