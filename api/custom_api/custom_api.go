package custom_api

import (
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/o/org/customize"
	"qrcode-bulk/qrcode-bulk-generator/x/web"
)

type CustomServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewCustomServer() *CustomServer {
	var s = &CustomServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/save", s.HandleSaveCus)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdateByQrID)
	s.HandleFunc("/getcusID", s.HandleGetByCusID)
	s.HandleFunc("/updatecusID", s.HandleUpdateByCusID)

	return s
}

func (s *CustomServer) HandleSaveCus(w http.ResponseWriter, r *http.Request) {
	var cus customize.Customize
	s.MustDecodeBody(r, &cus)
	web.AssertNil(cus.Create())
	s.SendData(w, cus)
}

func (s *CustomServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var cus = s.mustGetCustom(r)
	s.SendData(w, cus)
}

func (s *CustomServer) HandleGetByCusID(w http.ResponseWriter, r *http.Request) {
	var cus = s.mustGetCustomID(r)
	s.SendData(w, cus)
}

func (s *CustomServer) mustGetCustomID(r *http.Request) *customize.Customize {
	var qrID = r.URL.Query().Get("id")
	var cus, err = customize.GetByID(qrID)
	web.AssertNil(err)
	return cus
}

func (s *CustomServer) mustGetCustom(r *http.Request) *customize.Customize {
	var qrID = r.URL.Query().Get("qrcode_id")
	var cus, err = customize.GetByQrID(qrID)
	web.AssertNil(err)
	return cus
}

func (s *CustomServer) HandleUpdateByQrID(w http.ResponseWriter, r *http.Request) {
	var newCus customize.Customize
	s.MustDecodeBody(r, &newCus)
	var cus = s.mustGetCustom(r)
	web.AssertNil(cus.Update(&newCus))
	s.SendData(w, nil)
}

func (s *CustomServer) HandleUpdateByCusID(w http.ResponseWriter, r *http.Request) {
	var newCus customize.Customize
	s.MustDecodeBody(r, &newCus)
	var cus = s.mustGetCustomID(r)
	web.AssertNil(cus.Update(&newCus))
	s.SendData(w, nil)
}
