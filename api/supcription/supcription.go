package supcription

import (
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/api/auth/session"
	"qrcode-bulk/qrcode-bulk-generator/o/org/feature"
	"qrcode-bulk/qrcode-bulk-generator/o/org/supcription"
	"qrcode-bulk/qrcode-bulk-generator/o/org/user"
	"qrcode-bulk/qrcode-bulk-generator/x/web"
)

type SupcriptionServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewSupcriptionServer() *SupcriptionServer {
	var s = &SupcriptionServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/mark_delete", s.HandleMarkDelete)
	s.HandleFunc("/package", s.GetPackage)
	s.HandleFunc("/reg", s.Register)
	s.HandleFunc("/admin_reg", s.AdminRegister)
	return s
}

func (s *SupcriptionServer) AdminRegister(w http.ResponseWriter, r *http.Request) {
	var userID = r.URL.Query().Get("user_id")
	var pkgCode = r.URL.Query().Get("pkg_code")
	var pkg = feature.GetPkg(pkgCode)
	var sup = supcription.Supcription{}
	sup.PackageCode = pkgCode
	sup.QrcodeCount = pkg.QrcodeLimit

	web.AssertNil(sup.Create())

	var u, err = user.GetByID(userID)
	if err != nil {
		s.SendError(w, err)
	}
	u.SupcriptionID = sup.ID
	u.Update(u)

	s.SendData(w, map[string]interface{}{
		"user": u,
	})
}

func (s *SupcriptionServer) Register(w http.ResponseWriter, r *http.Request) {
	var us = session.MustAuthScope(r)

	var pkgCode = r.URL.Query().Get("pkg_code")
	var pkg = feature.GetPkg(pkgCode)
	var sup = supcription.Supcription{}
	sup.PackageCode = pkgCode
	sup.QrcodeCount = pkg.QrcodeLimit

	web.AssertNil(sup.Create())

	var u, err = user.GetByID(us.UserID)
	if err != nil {
		s.SendError(w, err)
	}

	u.SupcriptionID = sup.ID
	u.Update(u)

	s.SendData(w, map[string]interface{}{
		"user": u,
	})
}

func (s *SupcriptionServer) GetPackage(w http.ResponseWriter, r *http.Request) {
	s.SendData(w, feature.Pkgs)
}

func (s *SupcriptionServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var u supcription.Supcription
	s.MustDecodeBody(r, &u)
	web.AssertNil(u.Create())
	s.SendData(w, u)
	//cu.OnUserCreated(u.ID)
}

func (s *SupcriptionServer) mustGetSupcription(r *http.Request) *supcription.Supcription {
	var id = r.URL.Query().Get("id")
	var u, err = supcription.GetByID(id)
	web.AssertNil(err)
	return u
}

func (s *SupcriptionServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newUser supcription.Supcription
	s.MustDecodeBody(r, &newUser)
	var u = s.mustGetSupcription(r)
	web.AssertNil(u.Update(&newUser))
	s.SendData(w, nil)
}

func (s *SupcriptionServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetSupcription(r)
	s.SendData(w, u)
}

func (s *SupcriptionServer) HandleMarkDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetSupcription(r)
	web.AssertNil(supcription.MarkDelete(u.ID))
	s.Success(w)
}

// func (s *SupcriptionServer) HandleAllUser(w http.ResponseWriter, r *http.Request) {
// 	var res, err = supcription.GetAll()
// 	if err != nil {
// 		s.SendError(w, err)
// 	} else {
// 		s.SendData(w, res)
// 	}
// }
