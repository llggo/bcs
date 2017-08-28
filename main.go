package main

import (
	// other packages
	"os"
	"os/signal"

	"github.com/golang/glog"

	// 1. Config
	_ "qrcode-bulk/qrcode-bulk-generator/config"

	"qrcode-bulk/qrcode-bulk-generator/httpserver"
	"qrcode-bulk/qrcode-bulk-generator/x/runtime"
)

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	runtime.MaxProc()
	httpserver.Run(stop)
	glog.Flush()
}
