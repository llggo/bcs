package main

import (
	// other packages
	"os"
	"os/signal"

	"github.com/golang/glog"

	// 1. Config
	_ "bar-code/bcs/config"

	"bar-code/bcs/httpserver"
	"bar-code/bcs/x/runtime"
)

stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	runtime.MaxProc()
	httpserver.Run(stop)
	glog.Flush()
}
