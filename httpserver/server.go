package httpserver

import (
	"net/http"
	"qrcode-bulk/qrcode-bulk-generator/api"
	"qrcode-bulk/qrcode-bulk-generator/config"
	"qrcode-bulk/qrcode-bulk-generator/x/upload"
)

func serverHandler() http.Handler {
	var server = http.NewServeMux()

	// server.Handle("/", http.RedirectHandler("/app/", http.StatusFound))
	var app = newStatic(config.Station().Static.AppFolder)

	server.Handle("/app/", http.StripPrefix("/app/", app))

	var device = newStatic(config.Station().Static.DeviceFolder)
	server.Handle("/device/", http.StripPrefix("/device/", device))

	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	var up = upload.NewUploadFileServer("static/upload", 40960000)
	server.Handle("/upload/", http.StripPrefix("/upload/", up))

	// application specific
	server.Handle("/api/", http.StripPrefix("/api", api.NewApiServer()))

	return server
}
