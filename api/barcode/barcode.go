package barcode

import (
	"bar-code/bcs/o/org/barcode"
	"bar-code/bcs/x/web"
	"net/http"
)

type BarServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewBarServer() *BarServer {
	var s = &BarServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/find", s.HandleFindByName)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/mark_delete", s.HandleMarkDelete)
	s.HandleFunc("/search", s.HandleAllCode)
	return s
}

func (s *BarServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var u barcode.BarCode
	s.MustDecodeBody(r, &u)
	web.AssertNil(u.Create())
	s.SendData(w, u)
	//cu.OnUserCreated(u.ID)
}

func (s *BarServer) mustGetCode(r *http.Request) *barcode.BarCode {
	var id = r.URL.Query().Get("id")
	var u, err = barcode.GetByID(id)
	web.AssertNil(err)
	return u
}

func (s *BarServer) mustFindCode(r *http.Request) *barcode.BarCode {
	var code = r.URL.Query().Get("name")
	var c, err = barcode.GetByCode(code)
	web.AssertNil(err)
	return c
}

func (s *BarServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newCode barcode.BarCode
	s.MustDecodeBody(r, &newCode)
	var u = s.mustGetCode(r)

	web.AssertNil(u.Update(&newCode))

	s.SendData(w, nil)
}

func (s *BarServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetCode(r)
	s.SendData(w, u)
}

func (s *BarServer) HandleFindByName(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetCode(r)
	s.SendData(w, u)
}

func (s *BarServer) HandleMarkDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetCode(r)
	web.AssertNil(barcode.MarkDelete(u.ID))
	s.Success(w)
}

func (s *BarServer) HandleAllCode(w http.ResponseWriter, r *http.Request) {
	var res, err = barcode.GetAll()
	if err != nil {
		s.SendError(w, err)
	} else {
		s.SendData(w, res)
	}
}
