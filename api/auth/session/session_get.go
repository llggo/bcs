package session

import (
	"bar-code/bcs/o/auth/session"
	"bar-code/bcs/x/web"
)

const (
	errReadSessonFailed   = web.InternalServerError("read session failed")
	errSessionNotFound    = web.Unauthorized("session not found")
	errUnauthorizedAccess = web.Unauthorized("unauthorized access")
)

func Get(sessionID string) (*session.Session, error) {
	var s, err = session.GetByID(sessionID)
	if err != nil {
		sessionLog.Error(err)
		return nil, errReadSessonFailed
	}
	if s == nil {
		return nil, errSessionNotFound
	}
	return s, nil
}
