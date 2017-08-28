package socket

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type WsClient struct {
	UID    string          //UID: Socket Id
	Auth   Auth            //
	Socket *websocket.Conn //
}

func newWsClient(a Auth, s *websocket.Conn) *WsClient {
	var c = &WsClient{
		Auth:   a,
		Socket: s,
	}
	c.UID = fmt.Sprintf("%s", c)
	return c
}

func (c *WsClient) Close() {
	if c.Socket != nil {
		c.Socket.Close()
		// close(c.reply)
		c.Socket = nil
	}
}

func (c *WsClient) write(data []byte) {
	sendToReplyQueue(c, data)
}

func (c *WsClient) WriteError(err error) {
	c.write(BuildErrorMessage("/server", err))
}

func (c *WsClient) WriteJson(uri string, v interface{}) {
	c.write(BuildJsonMessage(uri, v))
}
