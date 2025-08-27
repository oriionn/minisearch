package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type PackagesResponseResult struct {
	Name         string `json:"pkgname"`
	Repository   string `json:"repo"`
	Architecture string `json:"arch"`
 	Description  string `json:"pkgdesc"`
  	Packager     string `json:"packager"`
}

type PackagesResponse struct {
	Results []PackagesResponseResult `json:"results"`
}

func Packages(query string) ([]SearchResult, error) {
	fetchUrl := fmt.Sprintf("https://archlinux.org/packages/search/json/?q=%s", url.QueryEscape(query))

	res, err := http.Get(fetchUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Arch Packages responded with a %d status code", res.StatusCode)
	}

	var rawResults PackagesResponse
	err = json.NewDecoder(res.Body).Decode(&rawResults)
	if err != nil {
		return nil, err
	}

	results := []SearchResult{}
	for _, result := range rawResults.Results {
		title := fmt.Sprintf("%s - Arch Linux", result.Name)
		description := fmt.Sprintf("%s · %s · %s · %s", result.Description, result.Repository, result.Architecture, result.Packager)
		link := fmt.Sprintf("https://archlinux.org/packages/%s/%s/%s", result.Repository, result.Architecture, result.Name)

		results = append(results, SearchResult{
			Title: title,
			Description: description,
			Link: link,
			Domain: "archlinux.org",
		})
	}

	return results, nil
}
