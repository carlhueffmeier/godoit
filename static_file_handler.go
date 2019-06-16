package main

import (
	"net/http"
	"path/filepath"
)

// StaticFileHandler serves static files from "public" folder
type StaticFileHandler struct{}

func (StaticFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("public", r.URL.Path)
	http.ServeFile(w, r, filePath)
}
