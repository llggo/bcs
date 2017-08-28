package api

import (
	"net/http"

	"qrcode-bulk/qrcode-bulk-generator/api/auth"
	"qrcode-bulk/qrcode-bulk-generator/api/user"
	"qrcode-bulk/qrcode-bulk-generator/x/web"
)

type ApiServer struct {
	*http.ServeMux
	web.JsonServer
}

func NewApiServer() *ApiServer {
	var s = &ApiServer{
		ServeMux: http.NewServeMux(),
	}

	s.Handle("/user/", http.StripPrefix("/user", user.NewUserServer()))
	s.Handle("/auth/", http.StripPrefix("/auth", auth.NewAuthServer()))
	// s.Handle("/qrcode-api/", http.StripPrefix("/qrcode-api", qrcode_api.NewQrCodeServer()))
	// s.Handle("/handle/", http.StripPrefix("/handle", url_handle.NewURLHandleServer()))
	// s.Handle("/customize/", http.StripPrefix("/customize", custom_api.NewCustomServer()))
	// s.Handle("/supcription/", http.StripPrefix("/supcription", supcription.NewSupcriptionServer()))
	// s.Handle("/bulk/", http.StripPrefix("/bulk", bulk.NewBulkServer()))
	// s.Handle("/verify_code/", http.StripPrefix("/verify_code", verify_code.NewVerifyCodeServer()))
	return s
}

func (s *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
