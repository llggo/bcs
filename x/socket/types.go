package socket

type EventHandler func(uri string, v interface{})
type IBoxHandler func(r *Request)

type Auth interface {
	ID() string
}

type AuthBearer string

func (a AuthBearer) ID() string {
	return string(a)
}

var AuthOff = AuthBearer("")

type ReadWriter interface {
	Read() ([]byte, bool)
	Write([]byte)
}
