package barcode

import (
	"bar-code/bcs/o/model"

	"gopkg.in/mgo.v2/bson"
)

var TableBarcode = model.NewTable("barcode", "bar")

func NewBarID() string {
	return TableBarcode.Next()
}

func AllUserOnBranch(branchid string, role string) ([]BarCode, error) {
	var res = make([]BarCode, 0)
	var query = bson.M{"branch_id": branchid}
	if role != "" {
		query["role"] = role
	}
	return res, TableBarcode.C().Find(query).All(&res)
}

func (b *BarCode) Create() error {
	return TableBarcode.Create(b)
}

func MarkDelete(id string) error {
	return TableBarcode.MarkDelete(id)
}

func (v *BarCode) Update(newValue *BarCode) error {
	var values = map[string]interface{}{
		"name": newValue.Name,
	}

	if newValue.GetName() != v.GetName() {
		values["name"] = newValue.GetName()
	}
	if newValue.GetUserID() != v.GetUserID() {
		values["user_id"] = newValue.GetUserID()
	}
	if newValue.GetStatus() != v.GetStatus() {
		values["status"] = newValue.GetStatus()
	}
	if newValue.GetType() != v.GetType() {
		values["type"] = newValue.GetType()
	}
	if newValue.GetImage() != v.GetImage() {
		values["type"] = newValue.GetImage()
	}
	if newValue.GetCode() != v.GetCode() {
		values["code"] = newValue.GetCode()
	}

	return TableBarcode.UnsafeUpdateByID(v.ID, values)
}
