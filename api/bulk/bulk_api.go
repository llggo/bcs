package bulk

import (
	"net/http"
	"bar-code/bcs/api/auth/session"
	"bar-code/bcs/o/org/bulk"
	"bar-code/bcs/x/web"
)

type BulkServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewBulkServer() *BulkServer {
	var s = &BulkServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdateByID)
	s.HandleFunc("/mark_delete", s.HandleMarkDelete)
	s.HandleFunc("/search", s.HandleAllBulk)
	s.HandleFunc("/export_excel", s.HandleExportExcel)
	s.HandleFunc("/change_status", s.HandleChangeStatus)
	s.HandleFunc("/updateVerify", s.HandleUpdateVerify)
	return s
}

func (s *BulkServer) HandleExportExcel(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	xlsx, err := bulk.ExportExcel(id)
	if err != nil {
		s.SendError(w, err)
		return
	}

	xlsx.Save()

	if err != nil {
		s.SendError(w, err)
		return
	}

	url := xlsx.FileName + ".xlsx"
	s.SendData(w, url)
	// http.Redirect(w, r, url, http.StatusSeeOther)
}

func (s *BulkServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)

	var d bulk.BulkType

	s.MustDecodeBody(r, &d)

	d.Bulk.UserID = u.UserID

	bulk.BulkChan <- d
	s.SendData(w, d)

	// if t == 4 {
	// 	var d bulk.BDGQGVFile
	// 	s.MustDecodeBody(r, &d)
	// 	d.Bulk.UserID = u.UserID
	// 	bulk.ChanCBGQGVFile <- d
	// 	s.SendData(w, d)
	// }
}

func (s *BulkServer) mustGetBulk(r *http.Request) *bulk.Bulk {
	var id = r.URL.Query().Get("id")
	var b, err = bulk.GetByID(id)
	web.AssertNil(err)
	return b
}

func (s *BulkServer) HandleUpdateByID(w http.ResponseWriter, r *http.Request) {
	var newBulk bulk.Bulk
	s.MustDecodeBody(r, &newBulk)
	var b = s.mustGetBulk(r)

	web.AssertNil(b.Update(&newBulk))

	s.SendData(w, nil)
}

func (s *BulkServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var b = s.mustGetBulk(r)
	s.SendData(w, b)
}

func (s *BulkServer) HandleMarkDelete(w http.ResponseWriter, r *http.Request) {
	var b = s.mustGetBulk(r)
	web.AssertNil(bulk.MarkDelete(b.ID))
	s.Success(w)
}

func (s *BulkServer) HandleAllBulk(w http.ResponseWriter, r *http.Request) {
	var u = session.MustAuthScope(r)
	var res, err = bulk.GetAll(map[string]interface{}{
		"user_id": u.UserID,
		"dtime":   0,
	})
	if err != nil {
		s.SendError(w, err)
	} else {
		s.SendData(w, res)
	}
}

func (s *BulkServer) HandleChangeStatus(w http.ResponseWriter, r *http.Request) {
	var b = s.mustGetBulk(r)
	b.Status = !b.Status
	web.AssertNil(b.Update(b))
	s.SendData(w, nil)
}

func (s *BulkServer) HandleUpdateVerify(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	err := bulk.UpdateVerify(id)
	if err != nil {
		s.SendError(w, err)
		return
	}
	s.SendData(w, nil)
}
