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


	results, err := Google(query, false)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	second_results, err := Google(query, true)
	if err == nil {
		for _, result := range second_results {
			results = append(results, result)
		}
	}

	return results, nil
}
