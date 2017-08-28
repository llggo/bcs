package qrcode

import (
	"github.com/mitchellh/mapstructure"
)

func (qr *QrCode) HandleText() (string, error) {
	var t TPEY
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return t.Content, nil
}

func (qr *QrCode) HandleUrl() (string, error) {
	var t URL
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return t.URL, nil
}

// func (s *QrCodeServer) HandleUrls(w http.ResponseWriter, r *http.Request, qr *qrcode.QrCode) string {
// 	var t map[string]qrcode.URL
// 	err := mapstructure.Decode(qr.GetActiveData(), &t)
// 	if err != nil {
// 		s.SendError(w, err)
// 	}
// 	return t.URL
// }

func (qr *QrCode) HandleSms() (string, error) {
	var t MS
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return "SMSTO:" + t.SMSTo + ":" + t.Content, nil
}

func (qr *QrCode) HandleMms() (string, error) {
	var t MS
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return "MMSTO:" + t.SMSTo + ":" + t.Content, nil
}

func (qr *QrCode) HandlePhone() (string, error) {
	var t Phone
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return "tel:+84" + t.Number, nil
}

func (qr *QrCode) HandleEmail() (string, error) {
	var t TPEY
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return "mailto:" + t.Content, nil
}

func (qr *QrCode) HandleCalendar() (string, error) {
	var t VCalendar
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	headerBegin := "BEGIN:VCALENDAR VERSION:1.0 BEGIN:VEVENT CATEGORIES:" + t.Categpries
	status := " STATUS:" + t.Status
	dtStart := " DTSTART:" + t.DtStart
	dtEnd := " DTEND:" + t.DtEnd
	sumary := " SUMMARY:" + t.Summary
	description := " DESCRIPTION" + t.Desciption
	bottomEnd := " CLASS:PRIVATE END:VEVENT END:VCALENDAR"
	return headerBegin + status + dtStart + dtEnd + sumary + description + bottomEnd, nil
}

func (qr *QrCode) HandleEvent() (string, error) {
	var t VEVENT
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	headerBegin := "BEGIN:VEVENT "
	sumary := " SUMMARY:" + t.Summary
	dtStart := " DTSTART:" + t.DtStart
	dtEnd := " DTEND:" + t.DtEnd
	bottomEnd := "END:VEVENT"
	return headerBegin + sumary + dtStart + dtEnd + bottomEnd, nil
}

func (qr *QrCode) HandleGeo() (string, error) {
	var t GEO
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return "geo:" + t.Longitude + "," + t.Latitude + ",100", nil
}

func (qr *QrCode) HandleYoutube() (string, error) {
	var t TPEY
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return t.Content, nil
}

func (qr *QrCode) HandleWifi() (string, error) {
	var t Wifi
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return "WIFI:T:WEP;S:" + t.Name + ";P:" + t.Pass + ";;", nil
}

func (qr *QrCode) HandleImage() (string, error) {
	var t IPA
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return t.Path, nil
}
func (qr *QrCode) HandlePDF() (string, error) {
	var t IPA
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return t.Path, nil
}
func (qr *QrCode) HandleAudio() (string, error) {
	var t IPA
	err := mapstructure.Decode(qr.GetActiveData(), &t)
	if err != nil {
		return "", err
	}
	return t.Path, nil
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
