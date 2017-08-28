package qrcode

import (
	"errors"
	"qrcode-bulk/qrcode-bulk-generator/o/model"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

var TableQrcode = model.NewTable("qrcode", "qrcode")

func NewQrID() string {
	return TableQrcode.Next()
}

func (qr *QrCode) Create() error {
	qr.Enable = true
	qr.SetID(qr.MakeID())

	if qr.IsDynamic() {
		qr.CreateImage(qr.MakeDynamicLink()) //dynamic
	} else {
		var sl, err = qr.MakeStaticLink()
		if err != nil {
			return err
		}
		qr.CreateImage(sl) //static
	}

	return qr.Insert()
}

func (qr *QrCode) Insert() error {
	qr.BeforeCreate()

	if qr.GetID() == "" {
		qr.SetID(TableQrcode.IdMaker.Next())
	}

	return TableQrcode.UnsafeInsert(qr)
}

func (qr *QrCode) MakeID() string {
	return TableQrcode.IdMaker.Next()
}

func MarkDelete(id string) error {
	return TableQrcode.MarkDelete(id)
}
func (qr *QrCode) Update(newValue *QrCode) error {
	var values = map[string]interface{}{
		"name": newValue.QrName,
	}
	if newValue.GetName() != qr.GetName() {
		values["name"] = newValue.GetName()
	}
	if newValue.GetType() != qr.GetType() {
		values["type"] = newValue.QrType
	}
	if !reflect.DeepEqual(newValue.GetActiveData(), qr.GetActiveData()) {
		values["data"] = newValue.GetActiveData()
	}
	if newValue.GetActiveTemplate() != qr.GetActiveTemplate() {
		values["template"] = newValue.GetActiveTemplate()
	}
	if newValue.GetPathBase64() != qr.GetPathBase64() {
		values["path_base"] = newValue.GetPathBase64()
	}
	if newValue.GetTagetScan() != qr.GetTagetScan() {
		values["taget_scan"] = newValue.GetTagetScan()
	}

	values["enable"] = newValue.Enable

	return TableQrcode.UnsafeUpdateByID(qr.ID, values)
}

func (qr *QrCode) Customize(newValue *QrCode) error {
	var values = map[string]interface{}{
		"name": newValue.QrName,
	}
	return TableQrcode.UnsafeUpdateByID(qr.ID, values)
}

func AllCounterOnBranch(branchID string) ([]QrCode, error) {
	var res = make([]QrCode, 0)
	query := TableQrcode.C().Find(bson.M{"branch_id": branchID})
	iter := query.Iter()
	if nil == iter {
		return nil, errors.New("\tERROR: Counter_table.go::AllCounterOnBranch")
	}
	var c QrCode
	for iter.Next(&c) {
		res = append(res, c)
	}
	return res, nil
}
