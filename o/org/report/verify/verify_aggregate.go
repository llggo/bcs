package verify

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Aggregate struct {
	QrCodeID  string `bson:"qrcode_id" json:"qrcode_id,omitempty"`
	UserID    string `bson:"user_id" json:"user_id,omitempty"`
	CountScan int    `bson:"count_scan" json:"count_scan"`
}

func Compute(fillter []bson.M) ([]*Verify, error) {
	var res = []*Verify{}
	var err = VerifyTable.C().Pipe(fillter).All(&res)
	return res, err
}

func CountScan(fillter bson.M) *Aggregate {
	var q = VerifyTable.C().Find(fillter)
	c, err := q.Count()
	if err != nil {
		fmt.Print(err)
	}
	var agg Aggregate
	agg.CountScan = c
	return &agg
}

func Count(where interface{}) (int, error) {
	return VerifyTable.UnsafeCount(where)
}

func CountOfUser(qrcode_id string, user_id string, start, end time.Time) (int, error) {
	var res = []*Verify{}
	var err error
	res, err = Compute([]bson.M{
		{"$match": bson.M{
			"user_id": user_id,
			"rptime":  bson.M{"$gte": start, "$lt": end},
		},
		},
	})

	return len(res), err
}
func CountOfQrcode(qrcode_id string, start, end time.Time) (int, error) {
	var res = []*Verify{}
	var err error
	res, err = Compute([]bson.M{
		{"$match": bson.M{
			"qrcode_id": qrcode_id,
			"rptime":    bson.M{"$gte": start, "$lt": end},
		},
		},
	})

	return len(res), err
}

func Compute2(filter bson.M, groupByField string) []*Aggregate {
	var match = filter
	var group = bson.M{}
	var groupBy = []string{"ctime", groupByField}
	var id = bson.M{}
	for _, b := range groupBy {
		group[b] = bson.M{"$first": "$" + b}
		id[b] = "$" + b
	}
	group["_id"] = id

	var ops = map[string]string{
		"qrcode_id": "sum",
	}

	for field, opt := range ops {
		group[field] = bson.M{"$" + opt: "$" + field}
	}

	var pipeline = []bson.M{
		{"$match": match},
		{"$group": group},
	}

	var res = []*Aggregate{}
	var err = VerifyTable.C().Pipe(pipeline).All(&res)
	if err != nil {
		fmt.Println(err)
	}
	return res
}
