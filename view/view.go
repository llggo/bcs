package view

import (
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/x/web"
)

type ViewServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewViewServer() *ViewServer {

	var s = &ViewServer{
		ServeMux: http.NewServeMux(),
	}
	var tpm = http.FileServer(http.Dir("static"))
	s.Handle("/", http.StripPrefix("/", tpm))
	return s
}

func (s *ViewServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer s.Recover(w)
	header := w.Header()
	header.Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	header.Add(
		"Access-Control-Allow-Methods",
		"OPTIONS, HEAD, GET, POST, DELETE",
	)
	header.Add(
		"Access-Control-Allow-Headers",
		"Content-Type, Content-Range, Content-Disposition",
	)
	header.Add(
		"Access-Control-Allow-Credentials",
		"true",
	)
	header.Add(
		"Access-Control-Max-Age",
		"2520000", // 30 days
	)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	s.ServeMux.ServeHTTP(w, r)
}
