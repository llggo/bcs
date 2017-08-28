package url_handle

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"qrcode-bulk/qrcode-bulk-generator/o/org/bulk"
	"qrcode-bulk/qrcode-bulk-generator/o/org/qrcode"
	"qrcode-bulk/qrcode-bulk-generator/o/org/report/scan_log"
	"qrcode-bulk/qrcode-bulk-generator/o/org/report/verify"
	"qrcode-bulk/qrcode-bulk-generator/o/org/verify_code"
	"qrcode-bulk/qrcode-bulk-generator/x/web"
	"time"

	"github.com/mitchellh/mapstructure"
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
	s.HandleFunc("/view", s.Handle)
	s.HandleFunc("/type", s.HandleType)
	s.HandleFunc("/welcome", s.HandleWelcome)
	s.HandleFunc("/verify_success", s.HandleVerifySuccess)
	s.HandleFunc("/product_info", s.HandleProductInfo)
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
		"qrcode_id":     qrcodeID,
		"product_name":  b.Name,
		"product_image": b.Product.Image,
		"price":         b.Product.Price,
		"hello":         b.Hello,
		"bulk_id":       b.ID,
	}

	if b.Type == 1 {
		data["show_verify"] = false
	} else {
		data["show_verify"] = true
	}

	temp.Execute(w, data)
}

func (s *URLHandleServer) HandleType(w http.ResponseWriter, r *http.Request) {

	var qrcodeID = r.URL.Query().Get("qrcode_id")

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

	//data
	var data = map[string]interface{}{}
	v, err := verify.GetByFirstTime(map[string]interface{}{
		"qrcode_id": qrcodeID,
	})

	if err != nil {
		data["verify_location"] = "Không có thông tin"
		data["verify_time"] = "Không có thông tin"
	} else {
		data["verify_location"] = fmt.Sprintf("%v-%v", v.LocationInfo.City, v.LocationInfo.Country)
		data["verify_time"] = time.Unix(v.MTime, 0)
	}

	vc, err := verify.Count(map[string]interface{}{
		"qrcode_id": qrcodeID,
	})
	if err != nil {
		data["verify_count"] = "Không có thông tin"
	} else {
		data["verify_count"] = vc
	}

	sc, err := scan_log.Count(map[string]interface{}{
		"qrcode_id": qrcodeID,
	})

	if err != nil {
		data["scan_count"] = "Không có thông tin"
	} else {
		data["scan_count"] = sc
	}

	if r.Method == "GET" {
		data["qrcode_id"] = qrcodeID

		//template
		var t *template.Template

		switch b.Type {
		case 2:
			t, err = renderTemplate("verify", "code_input")
			if err != nil {
				s.SendError(w, err)
				return
			}
		case 3:
			t, err = renderTemplate("verify", "code_show")
			if err != nil {
				s.SendError(w, err)
				return
			}

			v, err := verify_code.Find(map[string]interface{}{
				"qrcode_id": qrcodeID,
			})
			if err != nil {
				s.SendError(w, err)
				return
			}

			data["verify_code"] = v.VCode

		default:
			return
		}
		t.Execute(w, data)
	}

	if r.Method == "POST" {

		data["qrcode_id"] = r.FormValue("qrcodeID")

		switch b.Type {
		case 2:

			data["verify_code"] = r.FormValue("verifyCode")
			data["error"] = ""
			vc, err := verify_code.Find(map[string]interface{}{
				"qrcode_id": data["qrcode_id"],
				"code":      data["verify_code"],
			})
			//template
			var t *template.Template
			if err != nil {
				t, err = renderTemplate("verify", "code_input")
				if err != nil {
					s.SendError(w, err)
					return
				}
				data["error"] = "Invaild verify code"
				t.Execute(w, data)
				return
			}
			data["verify_code"] = vc.VCode
		case 3:

		default:
			return
		}

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

		url := fmt.Sprintf("/api/handle/verify_success?qrcode_id=%v&bulk_id=%v", data["qrcode_id"], b.ID)
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}

func (s *URLHandleServer) HandleVerifySuccess(w http.ResponseWriter, r *http.Request) {
	b, err := bulk.GetByID(r.URL.Query().Get("bulk_id"))
	if err != nil {
		s.SendError(w, err)
		return
	}
	t, err := renderTemplate("verify", "success")
	if err != nil {
		s.SendError(w, err)
		return
	}
	data := map[string]interface{}{
		"qrcode_id":     r.URL.Query().Get("qrcode_id"),
		"bulk_id":       b.ID,
		"price":         b.Product.Price,
		"product_image": b.Product.Image,
	}
	t.Execute(w, data)
}

func (s *URLHandleServer) HandleProductInfo(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{}

	var bulkID = r.URL.Query().Get("bulk_id")
	b, err := bulk.GetByID(bulkID)
	if err != nil {
		s.SendError(w, err)
		return
	}

	data["product"] = b.Product

	t, err := renderTemplate("verify", "product_info")
	if err != nil {
		s.SendError(w, err)
		return
	}
	t.Execute(w, data)
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

func (s *URLHandleServer) Handle(w http.ResponseWriter, r *http.Request) {
	var qrcodeID = r.URL.Query().Get("id")

	_, err := qrcode.GetByID(qrcodeID)
	if err != nil {
		s.SendError(w, err)
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		s.SendError(w, err)
		return
	}

	var sc scan_log.Scan
	sc.QrCodeID = qrcodeID
	sc.ScanIP = ip
	sc.DeviceInfo = scan_log.GetDeviceInfo(r)
	sc.Create()
	scan_log.SendScan(&sc)

	url := fmt.Sprintf("/api/handle/welcome?qrcode_id=%v", qrcodeID)
	http.Redirect(w, r, url, http.StatusSeeOther)

	// var id = r.URL.Query().Get("id")
	// var verifyID = r.URL.Query().Get("verify_id")
	// var urlError = "/api/handle/verify?qrcode_id=" + id

	// if verifyID != "" {
	// 	v, err := verify_code.GetByID(verifyID)
	// 	if err != nil {

	// 		http.Redirect(w, r, urlError, http.StatusSeeOther)
	// 		return
	// 	}
	// 	if v.QrcodeID != id {
	// 		http.Redirect(w, r, urlError, http.StatusSeeOther)
	// 		return
	// 	}
	// } else {
	// 	var url = "/api/handle/verify?qrcode_id=" + id
	// 	http.Redirect(w, r, url, http.StatusSeeOther)
	// 	return
	// }

	// //Get qrcode
	// var q, err = qrcode.GetByID(id)

	// if err != nil {
	// 	s.SendError(w, err)
	// 	return
	// }

	// //Get Bulk
	// b, err := bulk.GetByID(q.BulkID)

	// if err != nil {
	// 	s.SendError(w, err)
	// 	return
	// }

	// ip, _, err := net.SplitHostPort(r.RemoteAddr)

	// var sc scan_log.Scan
	// sc.QrCodeID = q.ID
	// sc.ScanIP = ip
	// sc.DeviceInfo = scan.GetDeviceInfo(r)
	// sc.UserID = q.UserID
	// sc.Create()

	// scan.SendScan(&sc)

	// if !q.Enable {
	// 	w.Write([]byte("<h1>QR Code này hết tiền nên không xem được nữa. Liên hệ <i>Hoàng</i> & <i>Tuấn Anh</i> để nạp thêm tiền</h1>"))
	// 	return
	// }

	// switch q.QrType {
	// case "text":
	// 	s.TextHandle(w, r, q, b)
	// case "url":
	// 	s.UrlHandle(w, r, q, b)
	// case "urls":
	// 	s.UrlsHandle(w, r, q, b)
	// case "image":
	// 	s.FileHandle(w, r, q, b)
	// case "pdf":
	// 	s.FileHandle(w, r, q, b)
	// case "audio":
	// 	s.FileHandle(w, r, q, b)
	// case "social":
	// 	s.SocialHandle(w, r, q, b)
	// }
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

func (s *URLHandleServer) TextHandle(w http.ResponseWriter, r *http.Request, q *qrcode.QrCode, b *bulk.Bulk) {

	var t qrcode.TPEY
	err := mapstructure.Decode(q.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}

	temp, err := renderTemplate(q.GetType(), q.GetActiveTemplate())

	if err != nil {
		s.SendError(w, err)
		return
	}

	data := map[string]interface{}{"data": t.Content, "bulk": b}

	temp.Execute(w, data)

}

func (s *URLHandleServer) UrlHandle(w http.ResponseWriter, r *http.Request, q *qrcode.QrCode, b *bulk.Bulk) {
	var t qrcode.URL
	err := mapstructure.Decode(q.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}

	if s.isValidUrl(t.URL, w) {
		http.Redirect(w, r, t.URL, http.StatusSeeOther)
	}
}

func (s *URLHandleServer) isValidUrl(raw string, w http.ResponseWriter) bool {
	_, err := url.ParseRequestURI(raw)
	if err != nil {
		s.SendError(w, err)
		return false
	}
	return true
}

func (s *URLHandleServer) UrlsHandle(w http.ResponseWriter, r *http.Request, q *qrcode.QrCode, b *bulk.Bulk) {
	var t map[string]qrcode.Urls
	err := mapstructure.Decode(q.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	var as []qrcode.Urls

	for _, v := range t {
		as = append(as, v)
	}
	temp, err := renderTemplate(q.GetType(), q.GetActiveTemplate())
	if err != nil {
		s.SendError(w, err)
		return
	}
	data := map[string][]qrcode.Urls{"data": as}
	temp.Execute(w, data)
}

func (s *URLHandleServer) SocialHandle(w http.ResponseWriter, r *http.Request, q *qrcode.QrCode, b *bulk.Bulk) {
	var t map[string]qrcode.URL
	err := mapstructure.Decode(q.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}

	// var as []qrcode.URL

	// for _, v := range t {
	// 	as = append(as, v)
	// }

	temp, err := renderTemplate(q.GetType(), q.GetActiveTemplate())
	if err != nil {
		s.SendError(w, err)
		return
	}

	for k, v := range t {
		if v.URL == "" {
			delete(t, k)
		}
	}

	data := map[string]map[string]qrcode.URL{"data": t}
	temp.Execute(w, data)
}

func (s *URLHandleServer) FileHandle(w http.ResponseWriter, r *http.Request, q *qrcode.QrCode, b *bulk.Bulk) {
	var t qrcode.File
	err := mapstructure.Decode(q.GetActiveData(), &t)
	if err != nil {
		s.SendError(w, err)
	}
	temp, err := renderTemplate(q.GetType(), q.GetActiveTemplate())
	if err != nil {
		s.SendError(w, err)
		return
	}
	data := map[string]qrcode.File{"data": t}
	temp.Execute(w, data)
}
