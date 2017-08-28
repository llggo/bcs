package socket

import (
	"qrcode-bulk/qrcode-bulk-generator/x/event/worker"

	"golang.org/x/net/websocket"
)

var replyWorker = worker.NewWorker()

func sendToReplyQueue(client *WsClient, data []byte) {
	if client == nil || client.Socket == nil {
		return
	}
	soc := client.Socket
	var job = func() {
		defer recover()
		websocket.Message.Send(soc, string(data))
	}
	replyWorker.Add(job)
}
