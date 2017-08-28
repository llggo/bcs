package scan_log

import (
	"qrcode-bulk/qrcode-bulk-generator/o/model"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var ScanTable = model.NewTable("scan_log", "scan_log")

func ForEach(filter bson.M, skip int, limit int, cb func(*Scan)) error {
	var q = ScanTable.C().Find(filter)
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
	var tr = &Scan{}
	for iter.Next(tr) {
		cb(tr)
	}
	return nil
}
func GetByID(id string) (*Scan, error) {
	var report Scan
	return &report, ScanTable.ReadByID(id, &report)
}

func (b *Scan) Create() error {
	if b.CTime.Unix() < 1 {
		b.CTime = time.Now()
	}
	return ScanTable.Create(b)
}

func (s *Scan) Update(newValue *Scan) error {
	return ScanTable.UnsafeUpdateByID(s.ID, &newValue)
}

func init() {
	if err := ScanTable.EnsureIndex("ctime"); err != nil {
		objreportLoging.Error("report index error", err)
	}
}
