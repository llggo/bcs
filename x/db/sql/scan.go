package sql

type RowWritable interface {
	WriteList() ([]interface{}, error)
}

type RowScannable interface {
	ScanList() []interface{}
}

type RowUpdatable interface {
	UpdateMap(v RowUpdatable) (map[string]interface{}, error)
}

type RowFactory func() RowScannable
