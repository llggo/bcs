package verify_code

import (
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/api/auth/session"
	"qrcode-bulk/qrcode-bulk-generator/o/org/qrcode"
	"qrcode-bulk/qrcode-bulk-generator/o/org/verify_code"
	"qrcode-bulk/qrcode-bulk-generator/x/math"
	"qrcode-bulk/qrcode-bulk-generator/x/web"
)

type VerifyCodeServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewVerifyCodeServer() *VerifyCodeServer {
	var s = &VerifyCodeServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/mark_delete", s.HandleMarkDelete)
	s.HandleFunc("/search", s.HandleAllBulk)

	return s
}

func (s *VerifyCodeServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var b verify_code.VerifyCode
	s.MustDecodeBody(r, &b)
	var id = r.URL.Query().Get("id")

	var qr, _ = qrcode.GetByID(id)
	b.QrcodeID = qr.GetID()
	b.VCode = math.RandString("", 12)
	web.AssertNil(b.Create())
	s.SendData(w, b)
}

func (s *VerifyCodeServer) mustGetCode(r *http.Request) *verify_code.VerifyCode {
	var id = r.URL.Query().Get("id")
	var b, err = verify_code.GetByID(id)
	web.AssertNil(err)
	return b
}

func (s *VerifyCodeServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newCode verify_code.VerifyCode
	s.MustDecodeBody(r, &newCode)
	var b = s.mustGetCode(r)

	web.AssertNil(b.Update(&newCode))

	s.SendData(w, nil)
}

func (s *VerifyCodeServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var b = s.mustGetCode(r)
	s.SendData(w, b)
}

func (s *VerifyCodeServer) HandleMarkDelete(w http.ResponseWriter, r *http.Request) {
	var b = s.mustGetCode(r)
	web.AssertNil(verify_code.MarkDelete(b.ID))
	s.Success(w)
}

func (s *VerifyCodeServer) HandleAllBulk(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)
	var res, err = verify_code.GetAll(map[string]interface{}{
		"user_id": u.UserID,
		"dtime":   0,
	})
	if err != nil {
		s.SendError(w, err)
	} else {
		s.SendData(w, res)
	}
}
