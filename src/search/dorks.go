package search

import (
	"fmt"
	"strings"
)

// Some websites based on https://github.com/Grafikart/grafisearch/blob/main/search/blocklist.go
var blocklist []string = []string{
	"pinterest.com",
	"allocine.com",
	"jeuxvideo.com",
	"pinterest.fr",
}

var wikilist []string = []string{
	"fr.wikipedia.org",
	"fr.wiktionary.org",
	"lgbtqia.fandom.com",
	"minecraft.wiki",
	"wiki.archlinux.org",
}

func AddDorks(query string) string {
	dorks := query
	for _, domain := range blocklist {
		dorks = fmt.Sprintf("%s -site:%s", dorks, domain)
	}

	if strings.Contains(query, "!wiki") {
		toReplace := "("

		for i, domain := range wikilist {
			or := "OR"
			if i >= len(wikilist) - 1 {
				or = ")"
			}

			toReplace = fmt.Sprintf("%s site:%s %s", toReplace, domain, or)
		}

		dorks = strings.ReplaceAll(dorks, "!wiki", toReplace)
	}
	return dorks
}
