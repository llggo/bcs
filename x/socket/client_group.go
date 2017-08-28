package socket

type WsClientByAuth map[string]*WsClient

func newWsClientByAuth() WsClientByAuth {
	return WsClientByAuth(map[string]*WsClient{})
}

func (r WsClientByAuth) empty() bool {
	return len(r) == 0
}

func (r WsClientByAuth) add(w *WsClient) {
	r[w.UID] = w
}

func (r WsClientByAuth) remove(w *WsClient) {
	delete(r, w.UID)
}

func (r WsClientByAuth) Send(payload []byte) {
	for _, w := range r {
		w.write(payload)
	}
}

//WsClientManager : [idAuth]
type WsClientManager map[string]WsClientByAuth

func NewWsClientManager() WsClientManager {
	return WsClientManager(map[string]WsClientByAuth{})
}

func (rb WsClientManager) ForEach(cb func(*WsClient)) {
	for _, byAuth := range rb {
		if byAuth != nil {
			for _, w := range byAuth {
				cb(w)
			}
		}
	}
}

func (rb WsClientManager) ForEachByAuth(id string, cb func(*WsClient)) {
	var byAuth = rb[id]
	if byAuth != nil {
		for _, w := range byAuth {
			cb(w)
		}
	}
}

func (rb WsClientManager) Add(w *WsClient, id string) {
	var r = rb[id]
	if r == nil {
		r = newWsClientByAuth()
		rb[id] = r
	}
	r.add(w)
}

func (rb WsClientManager) Remove(w *WsClient, id string) {
	var r = rb[id]
	if r == nil {
		return
	}
	r.remove(w)
	if r.empty() {
		delete(rb, id)
	}
}

func (rb WsClientManager) SendJson(uri string, v interface{}) {
	var payload = BuildJsonMessage(uri, v)
	//fmt.Println("Send Payload: ", string(payload))
	rb.SendRaw(payload)
}

func (rb WsClientManager) SendRaw(payload []byte) {
	for _, r := range rb {
		r.Send(payload)
	}
}

// SendToGroup send data to the group with auth id
// return the number of clients in the group
func (rb WsClientManager) SendToGroup(authID string, uri string, v interface{}) int {
	var r = rb[authID]
	if r == nil {
		return 0
	}
	var payload = BuildJsonMessage(uri, v)
	r.Send(payload)
	return len(r)
}

func (rb WsClientManager) Destroy() {
	rb.ForEach(func(w *WsClient) {
		w.Close()
	})
}
