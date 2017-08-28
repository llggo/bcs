package sql

import (
	"qrcode/pba/x/web"
)

const minCodeLength = 3
const maxCodeLength = 1024
const minNameLength = 6
const maxNameLength = 1024

const errCodeTooShort = web.BadRequest("Code must be at least 3 character")
const errNameTooShort = web.BadRequest("Name must be at least 6 character")

type CodeRow struct {
	Row
	Code string `json:"code"`
	Name string `json:"name"`
}

func (r *CodeRow) Columns() []string {
	return []string{"id", "mtime", "dtime", "code", "name"}
}

func (r *CodeRow) ScanList() []interface{} {
	return []interface{}{&r.ID, &r.MTime, &r.DTime, &r.Code, &r.Name}
}

func (r *CodeRow) WriteList(table *Table) ([]interface{}, error) {
	if len(r.Code) < minCodeLength {
		return nil, errCodeTooShort
	}

	var err = table.Unique(map[string]interface{}{
		"code":  r.Code,
		"dtime": 0,
	})
	if err != nil {
		return nil, err
	}

	if len(r.Name) < minNameLength {
		return nil, errNameTooShort
	}
	r.MTime = r.Now()
	return []interface{}{r.ID, r.MTime, r.DTime, r.Code, r.Name}, nil
}

func (r *CodeRow) UpdateMap(v *CodeRow, table *Table) (map[string]interface{}, error) {
	var where = map[string]interface{}{
		"mtime": r.Now(),
	}

	if r.Code != v.Code {
		if len(v.Code) < minCodeLength {
			return nil, errCodeTooShort
		}
		var err = table.Unique(map[string]interface{}{
			"code":  v.Code,
			"dtime": 0,
		})
		if err != nil {
			return nil, err
		}
		where["code"] = v.Code
	}

	if r.Name != v.Name {
		if len(v.Name) < minNameLength {
			return nil, errNameTooShort
		}
		where["name"] = v.Name
	}

	return where, nil
}

func (t *Table) Unique(where map[string]interface{}) error {
	var count, err = t.Count(where)
	if err != nil {
		return err
	}
	if count > 0 {
		return web.BadRequest("already exist")
	}
	return nil
}

type CodeRowBranch struct {
	RowBranch
	Code string `json:"code"`
	Name string `json:"name"`
}

func (r *CodeRowBranch) Columns() []string {
	return append(r.RowBranch.Columns(), "code", "name")
}

func (r *CodeRowBranch) ScanList() []interface{} {
	return append(r.RowBranch.ScanList(), &r.Code, &r.Name)
}

func (r *CodeRowBranch) WriteList(table *Table) ([]interface{}, error) {
	if len(r.Code) < minCodeLength {
		return nil, errCodeTooShort
	}

	var err = table.Unique(map[string]interface{}{
		"code":      r.Code,
		"branch_id": r.BranchID,
		"dtime":     0,
	})
	if err != nil {
		return nil, err
	}

	if len(r.Name) < minNameLength {
		return nil, errNameTooShort
	}

	r.MTime = r.Now()
	return append(r.RowBranch.WriteList(), r.Code, r.Name), nil
}

func (r *CodeRowBranch) UpdateMap(v *CodeRowBranch, table *Table) (map[string]interface{}, error) {
	var where = map[string]interface{}{
		"mtime": r.Now(),
	}

	if r.Code != v.Code {
		if len(v.Code) < minCodeLength {
			return nil, errCodeTooShort
		}
		var err = table.Unique(map[string]interface{}{
			"code":      v.Code,
			"branch_id": v.BranchID,
			"dtime":     0,
		})
		if err != nil {
			return nil, err
		}
		where["code"] = v.Code
	}

	if r.Name != v.Name {
		if len(v.Name) < minNameLength {
			return nil, errNameTooShort
		}
		where["name"] = v.Name
	}
	return where, nil
}
