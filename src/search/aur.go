package search

import (
	"cmp"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"
)

type AURResponseResult struct {
	Name         string  `json:"Name"`
 	Description  string  `json:"Description"`
  	Maintainer   string  `json:"Maintainer"`
   	Popularity 	 float64 `json:"Popularity"`
}

type AURResponse struct {
	Results []AURResponseResult `json:"results"`
}

func AUR(query string) ([]SearchResult, error) {
	fetchUrl := fmt.Sprintf("https://aur.archlinux.org/rpc/?v=5&type=search&arg=%s&by=name-desc", url.QueryEscape(query))

	res, err := http.Get(fetchUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Arch Linux User Respository responded with a %d status code", res.StatusCode)
	}

	var rawResults AURResponse
	err = json.NewDecoder(res.Body).Decode(&rawResults)
	if err != nil {
		return nil, err
	}

	slices.SortFunc(rawResults.Results, func (a, b AURResponseResult) int {
		return cmp.Compare(b.Popularity, a.Popularity)
	})

	results := []SearchResult{}
	for _, result := range rawResults.Results {
		title := fmt.Sprintf("%s - Arch Linux User Repository", result.Name)
		description := fmt.Sprintf("%s · %s", result.Description, result.Maintainer)
		link := fmt.Sprintf("https://aur.archlinux.org/packages/%s", result.Name)

		results = append(results, SearchResult{
			Title: title,
			Description: description,
			Link: link,
			Domain: "aur.archlinux.org",
		})
	}

	return results, nil
}
