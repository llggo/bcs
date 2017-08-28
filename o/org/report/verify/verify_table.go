package verify

import (
	"qrcode-bulk/qrcode-bulk-generator/o/model"

	"gopkg.in/mgo.v2/bson"
)

var VerifyTable = model.NewTable("verify_log", "verify")

func ForEach(filter bson.M, skip int, limit int, cb func(*Verify)) error {
	var q = VerifyTable.C().Find(filter)
	if skip > 0 {
		q.Skip(skip)
	}
	if limit > 0 {
		q.Limit(limit)
	}
	var iter = q.Iter()
	var err = iter.Err()
	if err != nil {
		return err
	}
	defer iter.Close()
	var tr = &Verify{}
	for iter.Next(tr) {
		cb(tr)
	}
	return nil
}

func GetByID(id string) (*Verify, error) {
	var report Verify
	return &report, VerifyTable.ReadByID(id, &report)
}

func GetByFirstTime(where interface{}) (*Verify, error) {
	var v Verify
	return &v, VerifyTable.UnsafeRunGetOneBySort(where, "ctime desc", &v)
}

func (b *Verify) Create() error {
	return VerifyTable.Create(b)
}

func (s *Verify) Update(newValue *Verify) error {
	return VerifyTable.UnsafeUpdateByID(s.ID, &newValue)
}

func init() {
	if err := VerifyTable.EnsureIndex("ctime"); err != nil {
		objreportLoging.Error("report index error", err)
	}
}
