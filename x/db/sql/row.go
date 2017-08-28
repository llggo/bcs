package sql

import (
	"time"
)

type Row struct {
	ID    string `json:"id"`
	MTime int32  `json:"mtime,omitempty"`
	DTime int32  `json:"dtime,omitempty"` // mark as deleted
}

func (r *Row) Columns() []string {
	return []string{"id", "mtime", "dtime"}
}

func (r *Row) ScanList() []interface{} {
	return []interface{}{&r.ID, &r.MTime, &r.DTime}
}

func (r *Row) WriteList() []interface{} {
	return []interface{}{r.ID, r.MTime, r.DTime}
}

func (r *Row) UpdateMap(v RowUpdatable) map[string]interface{} {
	var now = time.Now().Unix()
	return map[string]interface{}{
		"mtime": now,
	}
}

func (r *Row) Now() int32 {
	return int32(time.Now().Unix())
}

type RowBranch struct {
	Row
	BranchID string `json:"branch_id"`
	Settings MapSQL `json:"settings,omitempty"`
}

func (r *RowBranch) Columns() []string {
	return []string{"id", "mtime", "dtime", "branch_id", "settings"}
}

func (r *RowBranch) ScanList() []interface{} {
	return []interface{}{&r.ID, &r.MTime, &r.DTime, &r.BranchID, &r.Settings}
}

func (r *RowBranch) WriteList() []interface{} {
	return []interface{}{r.ID, r.MTime, r.DTime, r.BranchID, r.Settings}
}

func (r *RowBranch) UpdateMap(v RowUpdatable) map[string]interface{} {
	var now = time.Now().Unix()
	return map[string]interface{}{
		"mtime":     now,
		"branch_id": r.BranchID,
		"settings":  r.Settings,
	}
}
