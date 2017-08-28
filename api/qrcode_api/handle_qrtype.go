package qrcode_api

import (
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/o/org/qrcode"

	"github.com/mitchellh/mapstructure"
)

func (s *QrCodeServer) HandleText(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.TPEY
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return t.Content
}

func (s *QrCodeServer) HandleUrl(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.URL
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return t.URL
}

// func (s *QrCodeServer) HandleUrls(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
// 	var t map[string]qrcode.URL
// 	err := mapstructure.Decode(qr.GetActiveData(), &t)
// 	if err != nil {
// 		s.SendError(w, err)
// 	}
// 	return t.URL
// }

func (s *QrCodeServer) HandleSms(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.MS
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return "SMSTO:" + t.SMSTo + ":" + t.Content
}

func (s *QrCodeServer) HandleMms(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.MS
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return "MMSTO:" + t.SMSTo + ":" + t.Content
}

func (s *QrCodeServer) HandlePhone(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.Phone
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return "tel:+84" + t.Number
}

func (s *QrCodeServer) HandleEmail(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.TPEY
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return "mailto:" + t.Content
}

func (s *QrCodeServer) HandleCalendar(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.VCalendar
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	headerBegin := "BEGIN:VCALENDAR VERSION:1.0 BEGIN:VEVENT CATEGORIES:" + t.Categpries
	status := " STATUS:" + t.Status
	dtStart := " DTSTART:" + t.DtStart
	dtEnd := " DTEND:" + t.DtEnd
	sumary := " SUMMARY:" + t.Summary
	description := " DESCRIPTION" + t.Desciption
	bottomEnd := " CLASS:PRIVATE END:VEVENT END:VCALENDAR"
	return headerBegin + status + dtStart + dtEnd + sumary + description + bottomEnd
}

func (s *QrCodeServer) HandleEvent(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.VEVENT
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	headerBegin := "BEGIN:VEVENT "
	sumary := " SUMMARY:" + t.Summary
	dtStart := " DTSTART:" + t.DtStart
	dtEnd := " DTEND:" + t.DtEnd
	bottomEnd := "END:VEVENT"
	return headerBegin + sumary + dtStart + dtEnd + bottomEnd
}

func (s *QrCodeServer) HandleGeo(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.GEO
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return "geo:" + t.Longitude + "," + t.Latitude + ",100"
}

func (s *QrCodeServer) HandleYoutube(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.TPEY
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return t.Content
}

func (s *QrCodeServer) HandleWifi(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.Wifi
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return "WIFI:T:WEP;S:" + t.Name + ";P:" + t.Pass + ";;"
}

func (s *QrCodeServer) HandleImage(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.IPA
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return t.Path
}
func (s *QrCodeServer) HandlePDF(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.IPA
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return t.Path
}
func (s *QrCodeServer) HandleAudio(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
	var t qrcode.IPA
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	return t.Path
}

// func (s *QrCodeServer) HandleCard(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
// 	var t qrcode.VCard
// 	err := mapstructure.Decode(qr.GetActiveData(), &t)
// 	if err != nil {
// 		s.SendError(w, err)
// 	}
// 	return t.Content
// }

//geo:40.71872,-73.98905,100
