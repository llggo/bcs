package qrcode_api

import (
	"fmt"
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/api/auth/session"
	"qrcode-bulk/qrcode-bulk-generator/o/org/feature"
	"qrcode-bulk/qrcode-bulk-generator/o/org/qrcode"
	"qrcode-bulk/qrcode-bulk-generator/x/web"
)

type QrCodeServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewQrCodeServer() *QrCodeServer {
	var s = &QrCodeServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreateQr)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/markdelete", s.HandleMarkDelete)
	s.HandleFunc("/search", s.HandleAllQrCode)
	s.HandleFunc("/count", s.HandleCountQrCode)
	s.HandleFunc("/enable", s.HandleEnableByID)
	s.HandleFunc("/disable", s.HandleDisableByID)
	return s
}

func (s *QrCodeServer) HandleCreateQr(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)
	var qr qrcode.QrCode
	s.MustDecodeBody(r, &qr)
	qr.UserID = u.UserID
	web.AssertNil(qr.Create())

	s.SendData(w, qr)
}

func (s *QrCodeServer) MakeStaticLink(w http.ResponseWriter, r *http.Request, typeQr string, qr *qrcode.QrCode) string {
	switch typeQr {
	default:
		return ""
	case "text":
		return s.HandleText(w, r, qr)
	case "url":
		return s.HandleUrl(w, r, qr)
	case "urls":
		return s.HandleUrl(w, r, qr)
	case "sms":
		return s.HandleSms(w, r, qr)
	case "mms":
		return s.HandleMms(w, r, qr)
	case "phone":
		return s.HandlePhone(w, r, qr)
	case "email":
		return s.HandleEmail(w, r, qr)
	case "calendar":
		return s.HandleCalendar(w, r, qr)
	case "geo":
		return s.HandleGeo(w, r, qr)
	case "youtube":
		return s.HandleYoutube(w, r, qr)
	case "event":
		return s.HandleEvent(w, r, qr)
	// case "card":
	// return s.HandleCard(w, r, qr)
	case "wifi":
		return s.HandleWifi(w, r, qr)
	case "image":
		return s.HandleImage(w, r, qr)
	case "pdf":
		return s.HandlePDF(w, r, qr)
	case "audio":
		return s.HandleAudio(w, r, qr)
	}
}

func (s *QrCodeServer) MakeDynamicLink(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	return "http://mirascan.vn:31000/api/handle/view?id=" + qr.GetID()
}

func (s *QrCodeServer) mustGetQrCode(r *http.Request) *qrcode.QrCode {
	var id = r.URL.Query().Get("id")
	var qr, err = qrcode.GetByID(id)
	web.AssertNil(err)
	return qr
}

func (s *QrCodeServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var qr = s.mustGetQrCode(r)
	s.SendData(w, qr)
}

func (s *QrCodeServer) HandleEnableByID(w http.ResponseWriter, r *http.Request) {
	var qr = s.mustGetQrCode(r)
	qr.Enable = true

	web.AssertNil(qr.Update(qr))

	s.SendData(w, nil)
}

func (s *QrCodeServer) HandleDisableByID(w http.ResponseWriter, r *http.Request) {
	var qr = s.mustGetQrCode(r)
	qr.Enable = false
	web.AssertNil(qr.Update(qr))
	s.SendData(w, nil)
}

func (s *QrCodeServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newQr qrcode.QrCode

	s.MustDecodeBody(r, &newQr)

	fmt.Println(newQr)

	var qr = s.mustGetQrCode(r)
	web.AssertNil(qr.Update(&newQr))
	s.SendData(w, nil)
}

func (s *QrCodeServer) HandleMarkDelete(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)
	var ok, mes = u.CheckAccess(feature.Qrcode, feature.Delete)
	if !ok {
		s.SendData(w, mes)
		return
	}
	var qr = s.mustGetQrCode(r)
	web.AssertNil(qrcode.MarkDelete(qr.ID))
	s.Success(w)
}
func (s *QrCodeServer) HandleAllQrCode(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)
	var ok, mes = u.CheckAccess(feature.Qrcode, feature.List)
	if !ok {
		s.SendData(w, mes)
		return
	}

	qrs, err := qrcode.GetAll(map[string]interface{}{
		"user_id": u.UserID,
		"dtime":   0,
	})

	if err != nil {
		s.SendError(w, err)
	} else {
		s.SendData(w, qrs)
	}
}

func (s *QrCodeServer) HandleCountQrCode(w http.ResponseWriter, r *http.Request) {

	var u = session.MustAuthScope(r)

	qrs, err := qrcode.Count(map[string]interface{}{
		"user_id": u.UserID,
		"dtime":   0,
	})

	if err != nil {
		s.SendError(w, err)
	} else {
		s.SendData(w, qrs)
	}
}

//APi edit noi dung :  static:  gen lai, dynamic: k can
//Api quan ly noi dung : chien dich, nhom, share, report, tag, ...
