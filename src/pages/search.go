package pages

import (
	_ "embed"
	"fmt"
	"html/template"
	"minisearch/src/search"
	"minisearch/src/utils"
	"net/http"
	"os"
	"strings"
)

//go:embed templates/search.html
var searchContent string

type SearchPageData struct {
	Query string
	Results []search.SearchResult
	Calculation bool
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

	query := r.URL.Query().Get("q")

	if strings.Contains(query, "!google") {
		query = strings.ReplaceAll(query, "!google", "")
		query = search.AddDorks(query)
		http.Redirect(w, r, fmt.Sprintf("https://www.google.com/search?q=%s", query), 302)
		return
	}

	results, err := search.Search(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t, err := template.New("search").Parse(searchContent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := SearchPageData{
		Query: query,
		Results: results,
		Calculation: search.IsCalculation(query),
	}

	t.Execute(w, data)
}
