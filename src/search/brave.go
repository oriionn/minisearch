package search

import (
	"fmt"
	"io"
	"minisearch/src/utils"

	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func Brave(query string, second_page bool) ([]SearchResult, error) {
	start := 0
	if second_page {
		start = 1
		return nil, nil
	}

	// q := AddDorks(query)
	fetchUrl := fmt.Sprintf("https://search.brave.com/search?q=%s&offset=%d", url.QueryEscape(query), start)
	req, err := http.NewRequest(http.MethodGet, fetchUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("accept-language", "fr,fr-FR;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}


	if res.StatusCode != 200 {
		return nil, fmt.Errorf("DDG responded with a %d status code", res.StatusCode)
	}

	var body io.Reader
	body = res.Body
	if strings.Contains(res.Header.Get("Content-Type"), "ISO-8859") {
		decoder := charmap.ISO8859_1.NewDecoder()
		body = transform.NewReader(res.Body, decoder)
	}

	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	results := []SearchResult{}
	document.Find("div.snippet:not(.standalone)").Each(func(i int, item *goquery.Selection) {
		title := item.Find(".title").First()
		a := item.Find("a")
		link := a.AttrOr("href", "")
		link = strings.ReplaceAll(strings.Split(link, "&")[0], "/url?q=", "")
		unescaped, err := url.QueryUnescape(link)
		if err == nil {
			link = unescaped
		}
		description := item.Find(".snippet-description").Text()

		isValid, u := utils.IsValidURL(link)

		if title.Text() != "" && isValid {
			results = append(results, SearchResult{
				Title: title.Text(),
				Description: description,
				Link: link,
				Domain: u.Hostname(),
			})
		}
	})

	return results, nil
}

var descriptionWordList []string = []string{
	"Avis",
	"En stock",
}

func IsValidDescription(description string) bool {
	for _, word := range descriptionWordList {
		if strings.Contains(description, word) {
			return false
		}
	}

	pattern := strings.Repeat(" ", 5)
	matched, _ := regexp.MatchString(pattern, strings.TrimSpace(description))
	return !matched
}
