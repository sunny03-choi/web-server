package main

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func main() {
	http.Handle("/", new(staticHandler))
	http.ListenAndServe(":6666", nil)
}

type staticHandler struct {
	http.Handler
}

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	localPath := "wwwroot" + req.URL.Path
	content, err := ioutil.ReadFile(localPath)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(http.StatusText(404)))
	}
	contentType := getContentType(localPath)
	w.Header().Add("content-Type", contentType)
	w.Write(content)
}
func getContentType(localPath string) string {
	var contentType string
	ext := filepath.Ext(localPath)
	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case "png":
		contentType = "image/jpeg"
	default:
		contentType = "text/plain"
	}
	return contentType
}
