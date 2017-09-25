package auth

import (
	"net/http"
	"bar-code/bcs/api/auth/session"
	"bar-code/bcs/o/org/user"
)

func (s *AuthServer) handleMySettings(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)
	me, _ := user.GetByID(u.UserID)
	s.SendData(w, map[string]interface{}{
		"me": me,
	})
}
