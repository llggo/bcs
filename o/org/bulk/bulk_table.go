package bulk

import (
	"errors"
	"bar-code/bcs/o/model"

	"gopkg.in/mgo.v2/bson"
)

var TableBulk = model.NewTable("bulk", "bulk")

func NewBulkID() string {
	return TableBulk.Next()
}

func (b *Bulk) Create() error {
	return TableBulk.Create(b)
}

func (b *Bulk) Insert() error {
	b.BeforeCreate()

	if b.GetID() == "" {
		b.SetID(TableBulk.IdMaker.Next())
	}

	return TableBulk.UnsafeInsert(b)
}

func (b *Bulk) MakeID() string {
	return TableBulk.IdMaker.Next()
}

func MarkDelete(id string) error {
	return TableBulk.MarkDelete(id)
}
func (b *Bulk) Update(newValue *Bulk) error {
	var values = map[string]interface{}{
		"name": newValue.Name,
	}
	if newValue.GetName() != b.GetName() {
		values["name"] = newValue.GetName()
	}

	if newValue.GetHello() != b.GetHello() {
		values["hello"] = newValue.GetHello()
	}

	if newValue.GetVerifyLimit() != b.GetVerifyLimit() {
		values["verify_limit"] = newValue.GetVerifyLimit()
	}

	if newValue.GetContact() != b.GetContact() {
		values["contact"] = newValue.GetContact()
	}

	// if !reflect.DeepEqual(newValue.GetInfoProduct(), b.GetInfoProduct()) {
	values["info_product"] = newValue.GetInfoProduct()
	values["refer_product"] = newValue.GetRefer()
	values["image"] = newValue.Image
	values["logo"] = newValue.Logo
	// values["type"] = newValue.Type
	// }

	values["status"] = newValue.GetStatus()

	return TableBulk.UnsafeUpdateByID(b.ID, values)
}

func AllCounterOnBranch(branchID string) ([]Bulk, error) {
	var res = make([]Bulk, 0)
	query := TableBulk.C().Find(bson.M{"branch_id": branchID})
	iter := query.Iter()
	if nil == iter {
		return nil, errors.New("\tERROR: Counter_table.go::AllCounterOnBranch")
	}
	var c Bulk
	for iter.Next(&c) {
		res = append(res, c)
	}
	return res, nil
}
