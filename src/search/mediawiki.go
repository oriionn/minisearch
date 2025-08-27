package search

import (
	"encoding/json"
	"fmt"
	"io"
	"minisearch/src/utils"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type WikiResponse struct {
	Query struct {
		Pages map[string]WikiResponsePage
	} `json:"query"`
}

type WikiResponsePage struct {
	Title string `json:"title"`
	Extract string `json:"extract"`
}

func Mediawiki(query string, domain string, isFromWikimedia bool) ([]SearchResult, error) {
	path := "/w"
	if !isFromWikimedia {
		path = ""
	}
	fetchUrl := fmt.Sprintf("https://%s%s/api.php?action=opensearch&search=%s", domain, path, url.QueryEscape(query))
	req, err := http.NewRequest(http.MethodGet, fetchUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("user-agent", utils.GetUserAgent())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%s responded with a %d status code", domain, res.StatusCode)
	}

	var rawResults []any
	err = json.NewDecoder(res.Body).Decode(&rawResults)
	if err != nil {
		return nil, err
	}

	if len(rawResults) != 4 {
		return nil, fmt.Errorf("unexpected %s response format", domain)
	}

	titlesAny, _ := rawResults[1].([]any)
	urlsAny, _ := rawResults[3].([]any)
	titles := make([]string, len(titlesAny))
	urls := make([]string, len(urlsAny))

	for i, v := range titlesAny {
		titles[i], _ = v.(string)
	}

	for i, v := range urlsAny {
		urls[i], _ = v.(string)
	}

	if isFromWikimedia {
		min := len(titles)
		if min >= 50 {
			min = 49
		}
		titles = titles[:min]

		fetchUrl = fmt.Sprintf("https://%s%s/api.php?action=query&prop=extracts|pageimages&exintro&explaintext&piprop=original&titles=%s&format=json", domain, path, url.QueryEscape(strings.Join(titles, "|")))
		req, err = http.NewRequest(http.MethodGet, fetchUrl, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("user-agent", utils.GetUserAgent())

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("%s responded with a %d status code", domain, resp.StatusCode)
		}

		var rawResultsFinal WikiResponse
		err = json.NewDecoder(resp.Body).Decode(&rawResultsFinal)
		if err != nil {
			return nil, err
		}

		sampleUrl := urlsAny[0].(string)
		parts := strings.Split(sampleUrl, "/")
		path = parts[len(parts) - 2]

		var results []SearchResult
		for _, page := range rawResultsFinal.Query.Pages {
			description := page.Extract
			if len(description) >= 300 {
				description = description[:300] + "..."
			}

			result := SearchResult{
				Title: page.Title,
				Description: description,
				Domain: domain,
				Link: fmt.Sprintf("https://%s/%s/%s", domain, path, url.QueryEscape(page.Title)),
			}
			results = append(results, result)
		}

		return results, nil
	} else {
		if len(urls) >= 15 {
			urls = urls[:15]
		}

		var results []SearchResult
		for i, url := range urls {
			resp, err := http.Get(url)

			if err == nil && resp.StatusCode == 200 {
				var body io.Reader
				body = resp.Body
				if strings.Contains(resp.Header.Get("Content-Type"), "ISO-8859") {
					decoder := charmap.ISO8859_1.NewDecoder()
					body = transform.NewReader(resp.Body, decoder)
				}

				document, err := goquery.NewDocumentFromReader(body)
				if err == nil {
					descriptionEl := document.Find(".mw-content-ltr > p").First()
					description := strings.TrimSpace(descriptionEl.Text())
					if len(description) >= 300 {
						description = description[:300] + "..."
					}

					if description != "" {
						results = append(results, SearchResult{
							Title: titles[i],
							Description: description,
							Link: url,
							Domain: domain,
						})
					}
				}
			}
		}

		return results, nil
	}
}
