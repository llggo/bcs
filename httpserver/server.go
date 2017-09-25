package httpserver

import (
	"net/http"
	"bar-code/bcs/api"
	"bar-code/bcs/config"
	"bar-code/bcs/view"
	"bar-code/bcs/x/upload"
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
	server.Handle("/", http.StripPrefix("/v", view.NewViewServer()))

	return server
}
