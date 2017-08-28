package user

import (
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/o/org/user"
	"qrcode-bulk/qrcode-bulk-generator/x/web"
)

type UserServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewUserServer() *UserServer {
	var s = &UserServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/find", s.HandleFindByName)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/mark_delete", s.HandleMarkDelete)
	s.HandleFunc("/search", s.HandleAllUser)
	return s
}

func (s *UserServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var u user.User
	s.MustDecodeBody(r, &u)
	web.AssertNil(u.Create())
	s.SendData(w, u)
	//cu.OnUserCreated(u.ID)
}

func (s *UserServer) mustGetUser(r *http.Request) *user.User {
	var id = r.URL.Query().Get("id")
	var u, err = user.GetByID(id)
	web.AssertNil(err)
	return u
}

func (s *UserServer) mustFindUser(r *http.Request) *user.User {
	var username = r.URL.Query().Get("username")
	var u, err = user.GetByUsername(username)
	web.AssertNil(err)
	return u
}

func (s *UserServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newUser user.User
	s.MustDecodeBody(r, &newUser)
	var u = s.mustGetUser(r)

	web.AssertNil(u.Update(&newUser))

	s.SendData(w, nil)
}

func (s *UserServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetUser(r)
	s.SendData(w, u)
}

func (s *UserServer) HandleFindByName(w http.ResponseWriter, r *http.Request) {
	var u = s.mustFindUser(r)
	s.SendData(w, u)
}

func (s *UserServer) HandleMarkDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetUser(r)
	web.AssertNil(user.MarkDelete(u.ID))
	s.Success(w)
}

func (s *UserServer) HandleAllUser(w http.ResponseWriter, r *http.Request) {
	var res, err = user.GetAll()
	if err != nil {
		s.SendError(w, err)
	} else {
		s.SendData(w, res)
	}
}
