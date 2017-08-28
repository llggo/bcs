package socket

import (
	"qrcode-bulk/qrcode-bulk-generator/x/timer"
	"time"
)

const aliveUri = "/__alive"
const reloadUri = "/reload"

var now int64

func init() {
	timer.NewTimer(func() {
		now = time.Now().Unix()
	}).Schedule(time.Second, time.Second)
}

func (b *Box) KeepAlive() {
	b.Broadcast(aliveUri, now)
}

func (b *Box) ReloadAll() {
	b.Broadcast(reloadUri, nil)
}

func (b *Box) Reload(id string) {
	b.Clients.SendToGroup(id, reloadUri, nil)
}
