package session

import (
	"bar-code/bcs/x/db/mgo"
	"encoding/json"
)

type Session struct {
	mgo.BaseModel `bson:",inline"`
	Username      string `json:"username"`
	UserID        string `json:"userid"`
	BranchID      string `json:"branch_id"`
	CTime         int64  `json:"ctime"`
}

func (a *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}
