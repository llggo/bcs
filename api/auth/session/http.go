package session

import (
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/o/auth/session"
	"qrcode-bulk/qrcode-bulk-generator/x/mlog"
)

var accessToken = "token"
var sessionLog = mlog.NewTagLog("session_log")

func MustGet(r *http.Request) *session.Session {
	var sessionID = r.URL.Query().Get(accessToken)
	var s, e = Get(sessionID)
	if e != nil {
		panic(e)
	}
	return s
}

func MustAuthScope(r *http.Request) *session.Session {
	var query = r.URL.Query()
	var sessionID = query.Get(accessToken)
	var s, e = Get(sessionID)
	if e != nil {
		panic(e)
	}
	return s
}

func MustClear(r *http.Request) {
	var sessionID = r.URL.Query().Get(accessToken)
	var e = session.MarkDelete(sessionID)
	if e != nil {
		sessionLog.Error(e, "remove session")
	}
}
