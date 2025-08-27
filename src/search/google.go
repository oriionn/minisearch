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

func Google(query string, second_page bool) ([]SearchResult, error) {
	start := 0
	if second_page {
		start = 10
	}

	q := AddDorks(query)
	fetchUrl := fmt.Sprintf("https://www.google.com/search?q=%s&start=%d", url.QueryEscape(q), start)
	req, err := http.NewRequest(http.MethodGet, fetchUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("accept-language", "fr,fr-FR;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("user-agent", utils.GetUserAgent())
	req.AddCookie(&http.Cookie{
		Name: "CONSENT",
		Value: "PENDING+987",
	})
	req.AddCookie(&http.Cookie{
		Name: "SOCS",
		Value: "CAESHAgBEhIaAB",
	})

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Google responded with a %d status code", res.StatusCode)
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
	document.Find(".ezO2md").Each(func(i int, item *goquery.Selection) {
		titleEl := item.Find(".CVA68e").First()
		var descriptionEl *goquery.Selection
		linkEl := item.Find("a").First()

		item.Find(".FrIlee").Each(func (_ int, ii *goquery.Selection) {
			if IsValidDescription(ii.Text()) && descriptionEl == nil {
				descriptionEl = ii
			}
		})

		if descriptionEl == nil {
			descriptionEl = item.Find(".FrIlee").First()
		}

		link := linkEl.AttrOr("href", "")
		link = strings.ReplaceAll(strings.Split(link, "&")[0], "/url?q=", "")
		unescaped, err := url.QueryUnescape(link)
		if err == nil {
			link = unescaped
		}

		title := strings.TrimSpace(titleEl.Text())
		description := strings.TrimSpace(descriptionEl.Text())
		validUrl, u := utils.IsValidURL(link)

		if title != "" && validUrl {
			results = append(results, SearchResult{
				Title: title,
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
