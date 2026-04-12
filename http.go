package goutils

import (
	"embed"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ServeCSS(r chi.Router, path, css string) {
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(css))
	})
}

func ServeJS(r chi.Router, path, js string) {
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(js))
	})
}

func ServeFS(r chi.Router, basePath string, fs embed.FS) {
	r.Handle(basePath+"*", http.StripPrefix(basePath, http.FileServer(http.FS(fs))))
}

func ServeFile(r chi.Router, path, content, contentType string) {
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(content))
	})
}
