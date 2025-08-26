package search

import "fmt"

// Some websites based on https://github.com/Grafikart/grafisearch/blob/main/search/blocklist.go
var blocklist []string = []string{
	"pinterest.com",
	"allocine.com",
	"jeuxvideo.com",
	"pinterest.fr",
}

func GenerateDorks() string {
	dorks := ""
	for _, domain := range blocklist {
		dorks = fmt.Sprintf("%s -site:%s", dorks, domain)
	}

	return dorks
}
