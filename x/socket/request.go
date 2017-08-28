package socket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type Request struct {
	Payload []byte
	RawURI  string
	URI     *url.URL
	Data    []byte
	Client  *WsClient
}

func NewRequest(w *WsClient, payload []byte) (*Request, error) {
	var r = &Request{
		Client:  w,
		Payload: payload,
	}

	var endOfURI = bytes.Index(payload, []byte(" "))
	var remaining = payload
	if endOfURI < 0 {
		r.RawURI = string(remaining)
		remaining = remaining[0:0]
	} else {
		r.RawURI = string(remaining[:endOfURI])
		remaining = remaining[endOfURI+1:]
	}

	var err error

	r.URI, err = url.ParseRequestURI(r.RawURI)
	if err != nil {
		return nil, err
	}
	r.Data = remaining
	return r, nil
}

func (r *Request) Path() string {
	return r.URI.Path
}

func (r *Request) UnmarshalJson(v interface{}) error {
	return WrapBadRequest(json.Unmarshal(r.Data, v), "unmarshal json")
}

func (r *Request) String() string {
	return fmt.Sprintf("url: [%s], data: [%s]\n", r.RawURI, r.Data)
}

func (r *Request) Reply(v interface{}) {
	r.Client.WriteJson(r.RawURI, v)
}

func (r *Request) ReplyError(err error) {
	r.Client.WriteJson("/error", map[string]interface{}{
		"uri": r.RawURI,
		"err": err.Error(),
	})
}

func (r *Request) isError() bool {
	return strings.HasPrefix(r.URI.Path, "/error")
}

func (r *Request) AuthID() string {
	return r.Client.Auth.ID()
}
