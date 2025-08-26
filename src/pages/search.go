package pages

import (
	_ "embed"
	"html/template"
	"minisearch/src/utils"
	"net/http"
	"os"
)

//go:embed templates/search.html
var searchContent string

type SearchPageData struct {
	Query string
}

func Search(w http.ResponseWriter, r *http.Request) {
	if (utils.DevMode()) {
		contentBytes, err := os.ReadFile("src/pages/templates/search.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		searchContent = string(contentBytes)
	}

	if !r.URL.Query().Has("q") {
		http.Error(w, "No query provided", http.StatusBadRequest)
	}

	t, err := template.New("search").Parse(searchContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := SearchPageData{
		Query: r.URL.Query().Get("q"),
	}

	t.Execute(w, data)
}
