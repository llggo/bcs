package session

import (
	"encoding/json"
	"qrcode-bulk/qrcode-bulk-generator/x/math"
)

var idMaker = math.RandStringMaker{Length: 40, Prefix: "s"}

type Session struct {
	SessionID string `json:"id"`
	Username  string `json:"username"`
	UserID    string `json:"user_id"`
	CTime     int64  `json:"ctime"`
}

func (a *Session) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}
