package pages

import (
	_ "embed"
	"io"
	"minisearch/src/utils"
	"net/http"
	"os"
)

//go:embed templates/index.html
var indexContent string

func Index(w http.ResponseWriter, r *http.Request) {
	if (utils.DevMode()) {
		contentBytes, err := os.ReadFile("src/pages/templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		indexContent = string(contentBytes)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, indexContent)
}
