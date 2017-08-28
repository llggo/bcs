package httpserver

import (
	"http/gziphandler"
	"net/http"
)

func newStatic(dir string) http.Handler {
	var withoutGzip = http.FileServer(http.Dir(dir))
	return gziphandler.GzipHandler(withoutGzip)
}
