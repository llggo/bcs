package auth

import (
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/api/auth/session"
	"qrcode-bulk/qrcode-bulk-generator/o/org/user"
	"qrcode-bulk/qrcode-bulk-generator/x/web"
	"strings"
)

type AuthServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewAuthServer() *AuthServer {
	var s = &AuthServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/login", s.HandleLogin)
	s.HandleFunc("/me", s.handleMe)
	s.HandleFunc("/my_settings", s.handleMySettings)
	s.HandleFunc("/logout", s.HandleLogout)
	s.HandleFunc("/change_pass", s.handleChangePass)
	return s
}

func (s *AuthServer) handleMe(w http.ResponseWriter, r *http.Request) {
	s.SendJson(w, session.MustGet(r))
}

func (s *AuthServer) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var body = struct {
		Username   string
		Password   string
		Scope      string
		BranchCode string `json:"branch_code"`
		Auto       bool
	}{}

	s.MustDecodeBody(r, &body)

	var u, err = user.GetByUsername(strings.ToLower(body.Username))
	if user.TableUser.IsErrNotFound(err) {
		s.SendError(w, errUserNotFound)
		return
	}

	web.AssertNil(err)

	if err = u.ComparePassword(body.Password); err != nil {
		s.SendError(w, err)
		return
	}

	if err != nil {
		s.SendError(w, err)
	}

	var ses = session.MustNew(u)

	w.Header().Set("X-QRCode-Token", ses.ID)
	s.SendData(w, map[string]interface{}{
		"user":    u,
		"session": ses,
	})
}

func (s *AuthServer) HandleLogout(w http.ResponseWriter, r *http.Request) {
	session.MustClear(r)
	s.SendData(w, nil)
}

func (s *AuthServer) handleChangePass(w http.ResponseWriter, r *http.Request) {
	// var u = session.MustGet(r)
	// me := user.GetByID(u.UserID)
	// pass := me.Password
	// fmt.Print(pass)
	// fmt.Print(me)
	// fmt.Print(u)
	var body = struct {
		OldPass   string `json:"old_pass"`
		NewPass   string `json:"new_pass"`
		ReNewPass string `json:"re_new_pass"`
		Username  string `json:"username"`
	}{}

	s.MustDecodeBody(r, &body)

	var u, err = user.GetByUsername(strings.ToLower(body.Username))
	if user.TableUser.IsErrNotFound(err) {
		s.SendError(w, errUserNotFound)
		return
	}

	if err := u.ComparePassword(body.OldPass); err != nil {
		s.SendError(w, err)
		return
	}
	// u.UpdatePass(body.NewPass)

	s.SendData(w, map[string]interface{}{
		"status": "success",
	})
}
