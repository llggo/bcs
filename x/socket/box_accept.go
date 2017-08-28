package socket

import (
	"golang.org/x/net/websocket"
	"runtime/debug"
)

func (b *Box) AcceptPublic(ws *websocket.Conn, args ...Auth) {
	if len(args) < 1 {
		b.AcceptDefault(ws, AuthOff)
	} else {
		b.AcceptDefault(ws, args[0])
	}
}

func (b *Box) AcceptDefault(ws *websocket.Conn, a Auth) {
	b.Accept(ws, a, b.Join, b.Leave)
}

func (b *Box) Accept(ws *websocket.Conn, a Auth, onJoin func(*WsClient), onLeave func(*WsClient)) {
	var codec = websocket.Message
	var c = newWsClient(a, ws)
	b.Clients.Add(c, c.Auth.ID())
	if onJoin != nil {
		onJoin(c)
	}

	defer func() {
		b.Clients.Remove(c, c.Auth.ID())
		if onLeave != nil {
			onLeave(c)
		}
		c.Close()
	}()

	for {
		var data []byte
		if err := codec.Receive(ws, &data); err != nil {
			break
		}
		var r, err = NewRequest(c, data)

		//fmt.Printf("ID: %s \tData: %s\n", b.ID, string(r.Payload))
		if err != nil {
			c.WriteError(err)
		} else {
			b.Serve(r)
		}
	}
}

var (
	errHandlerNotFound = BadRequest("HANDLER NOT FOUND")
	errInternalServer  = InternalServerError("SERVER ERROR")
)

func (b *Box) notFound(r *Request) {
	if r.isError() {
		boxLog.Errorln("Handler not found : " + r.RawURI)
	} else {
		r.ReplyError(errHandlerNotFound)
	}
}

func (b *Box) defaultRecover(r *Request, rc interface{}) {
	if err, ok := rc.(error); ok {
		if _, ok = err.(IWebError); ok {
			r.ReplyError(err)
			return
		}
		boxLog.Error(err, string(debug.Stack()))
		r.ReplyError(errInternalServer)
	} else {
		boxLog.Error(rc, string(debug.Stack()))
		r.ReplyError(errInternalServer)
	}
}

func (b *Box) join(w *WsClient) {

}

func (b *Box) leave(w *WsClient) {

}

func (b *Box) WriteError(ws *websocket.Conn, err error) {
	websocket.Message.Send(ws, string(BuildErrorMessage("/system", err)))
}
