package utils

import "net/url"

func IsValidURL(str string) (bool, *url.URL) {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != "", u
}
