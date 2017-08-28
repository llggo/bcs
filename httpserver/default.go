package httpserver

import (
	"net/http"
	"os"
	"qrcode-bulk/qrcode-bulk-generator/config"

	"github.com/golang/glog"
)

func Run(stop <-chan os.Signal) {
	var c = config.Station().Server
	var h = &http.Server{
		Addr:         c.Addr(),
		TLSNextProto: nil,
		Handler:      serverHandler(),
	}
	go func() {
		logger.Infof(0, "Listening on http://%s\n", c.Addr())
		if err := h.ListenAndServe(); err != nil {
			glog.Flush()
			glog.Exitf("Server %s", err.Error())
		}
	}()

	<-stop
	// logger.Infoln(0, "Shutting down the server...")
	// ctx, _ := c.Wait()
	// err := h.Shutdown(ctx)
	// if err == nil {
	// logger.Infoln(0, "Server gracefully stopped")
	// } else {
	// logger.Errorf("Server shutdown %s\n", err.Error())
	// }
}
