package verify_code

import (
	"errors"
	"qrcode-bulk/qrcode-bulk-generator/o/model"

	"gopkg.in/mgo.v2/bson"
)

var TableVerifyCode = model.NewTable("verify_code", "verify_code")

func NewBulkID() string {
	return TableVerifyCode.Next()
}

func (b *VerifyCode) Create() error {
	return TableVerifyCode.Create(b)
}

func (b *VerifyCode) Insert() error {
	b.BeforeCreate()

	if b.GetID() == "" {
		b.SetID(TableVerifyCode.IdMaker.Next())
	}

	return TableVerifyCode.UnsafeInsert(b)
}

func (b *VerifyCode) MakeID() string {
	return TableVerifyCode.IdMaker.Next()
}

func MarkDelete(id string) error {
	return TableVerifyCode.MarkDelete(id)
}
func (b *VerifyCode) Update(newValue *VerifyCode) error {
	var values = map[string]interface{}{
		"name": newValue.VName,
	}
	if newValue.GetName() != b.GetName() {
		values["name"] = newValue.GetName()
	}

	if newValue.GetCode() != b.GetCode() {
		values["code"] = newValue.GetCode()
	}

	return TableVerifyCode.UnsafeUpdateByID(b.ID, values)
}

func AllCounterOnBranch(branchID string) ([]VerifyCode, error) {
	var res = make([]VerifyCode, 0)
	query := TableVerifyCode.C().Find(bson.M{"branch_id": branchID})
	iter := query.Iter()
	if nil == iter {
		return nil, errors.New("\tERROR: Counter_table.go::AllCounterOnBranch")
	}
	var c VerifyCode
	for iter.Next(&c) {
		res = append(res, c)
	}
	return res, nil
}
