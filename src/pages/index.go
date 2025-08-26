package pages

import (
	"io"
	"net/http"
	_ "embed"
)

//go:embed templates/index.html
var indexContent string

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, indexContent)
}
