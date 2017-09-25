package url_handle

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"bar-code/bcs/o/org/bulk"
	"bar-code/bcs/o/org/qrcode"
	"bar-code/bcs/o/org/report/scan_log"
	"bar-code/bcs/o/org/report/verify"
	"bar-code/bcs/o/org/verify_code"
	"bar-code/bcs/x/web"
	"strconv"
	"time"
)

const view = "static/view"

type URLHandleServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewURLHandleServer() *URLHandleServer {
	var s = &URLHandleServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/welcome", s.HandleWelcome)
	s.HandleFunc("/verify", s.HandleVerify)
	s.HandleFunc("/get_template", s.TemplateByType)
	s.HandleFunc("/view_ip", s.HandleViewIP)
	return s
}

func (s *URLHandleServer) HandleWelcome(w http.ResponseWriter, r *http.Request) {
	var qrcodeID = r.URL.Query().Get("qrcode_id")
	temp, err := renderTemplate("verify", "welcome")
	if err != nil {
		s.SendError(w, err)
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		s.SendError(w, err)
		return
	}

	qr, err := qrcode.GetByID(qrcodeID)
	if err != nil {
		s.SendError(w, err)
		return
	}

	b, err := bulk.GetByID(qr.BulkID)
	if err != nil {
		s.SendError(w, err)
		return
	}

	var data = map[string]interface{}{
		"qrcode": qr,
		"bulk":   b,
	}

	if b.Type == 3 {
		v, err := verify_code.Find(map[string]interface{}{
			"qrcode_id": qrcodeID,
		})
		if err != nil {
			s.SendError(w, err)
			return
		}
		data["verify_code"] = v.VCode
	}

	if b.Type == 4 && qr.VerifyLimit > 1 {
		aw, err := scan_log.Find(map[string]interface{}{
			"qrcode_id": qrcodeID,
			"scan_ip":   ip,
		})
		if err != nil {
			qr.VerifyLimit--
			qr.Update(qr)
		} else {
			fmt.Println("ScanIP:", aw, time.Now())
		}
	}
	data["countlimit"] = b.VerifyLimit - qr.VerifyLimit

	v, err := verify.GetByFirstTime(map[string]interface{}{
		"qrcode_id": qrcodeID,
	})
	var sc scan_log.Scan
	sc.QrCodeID = qrcodeID
	sc.ScanIP = ip
	sc.DeviceInfo = scan_log.GetDeviceInfo(r)
	sc.Create()
	scan_log.SendScan(&sc)

	x, err := scan_log.GetByFirstTime(map[string]interface{}{
		"qrcode_id": qrcodeID,
	})
	if err == nil {
		if b.Type != 4 {
			data["active_location"] = v.LocationInfo
			data["active_time"] = time.Unix(v.MTime, 0)
		} else {
			data["active_location"] = x.LocationInfo
			data["active_time"] = time.Unix(x.MTime, 0)
		}
	}

	vc, err := verify.Count(map[string]interface{}{
		"qrcode_id": qrcodeID,
	})

	if err == nil {
		data["verify_count"] = vc
	}

	scount, err := scan_log.Count(map[string]interface{}{
		"qrcode_id": qrcodeID,
	})

	if err == nil {
		data["scan_count"] = scount
	}

	if b.Type == 1 {
		data["verify"] = false
	} else {
		data["verify"] = true
	}

	temp.Execute(w, data)
}

func (s *URLHandleServer) HandleVerify(w http.ResponseWriter, r *http.Request) {

	var data = map[string]interface{}{
		"status": true,
	}

	bulkType, err := strconv.ParseInt(r.URL.Query().Get("bulk_type"), 10, 64)

	if err != nil {
		data["status"] = false

	} else {
		var qrcodeID = r.URL.Query().Get("qrcode_id")
		var verifyCode = r.URL.Query().Get("verify_code")

		switch bulkType {
		case 4:
			err = nil
		default:
			_, err := qrcode.GetByID(qrcodeID)
			if err != nil {
				data["status"] = false
			}

			_, err = verify_code.Find(map[string]interface{}{
				"qrcode_id": qrcodeID,
				"code":      verifyCode,
			})
		}

		if err != nil {
			data["status"] = false
		} else {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				s.SendError(w, err)
				return
			}

			var v verify.Verify
			v.QrCodeID = qrcodeID
			v.ScanIP = ip
			v.DeviceInfo = verify.GetDeviceInfo(r)
			v.Create()
			verify.SendScan(&v)
		}
	}
	s.SendJson(w, data)
}

type Empty struct{}

func (s *URLHandleServer) HandleViewIP(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	s.SendError(w, err)
	s.SendData(w, map[string]interface{}{
		"ip":     ip,
		"device": verify.GetDeviceInfo(r),
		"loc":    verify.GetLocation(ip),
	})
}

func (s *URLHandleServer) TemplateByType(w http.ResponseWriter, r *http.Request) {
	var t = r.URL.Query().Get("type")

	temp, err := ioutil.ReadDir(view + "/" + t)

	if err != nil {
		s.SendError(w, err)
	}

	var as []string

	for _, d := range temp {
		if d.IsDir() {
			as = append(as, d.Name())
		}
	}

	s.SendData(w, as)
}
