package search

import (
	"strings"
)

func Search(query string) ([]SearchResult, error) {
	if strings.Contains(query, "!package") {
		query = strings.ReplaceAll(query, "!package", "")
		results, err := Packages(query)
		if err != nil {
			return nil, err
		}

		second_results, err := AUR(query)
		if err == nil {
			for _, result := range second_results {
				results = append(results, result)
			}
		}

		return results, nil
	}

	if strings.Contains(query, "!arch") {
		query = strings.ReplaceAll(query, "!arch", "")
		return Mediawiki(query, "wiki.archlinux.org", false)
	}

	if strings.Contains(query, "!wp") {
		query = strings.ReplaceAll(query, "!wp", "")
		return Mediawiki(query, "fr.wikipedia.org", true)
	}

	results, err := Brave(query, false)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	second_results, err := Brave(query, true)
	if err == nil {
		for _, result := range second_results {
			results = append(results, result)
		}
	}

	return results, nil
}
