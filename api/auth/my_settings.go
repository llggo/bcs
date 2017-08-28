package auth

import (
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/api/auth/session"
	"qrcode-bulk/qrcode-bulk-generator/o/org/user"
)

func (s *AuthServer) handleMySettings(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)
	me, _ := user.GetByID(u.UserID)
	s.SendData(w, map[string]interface{}{
		"me": me,
	})
}
