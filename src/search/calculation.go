package search

import "regexp"

func IsCalculation(query string) bool {
	re := regexp.MustCompile(`^[0-9+\-*/%^().\s]+$`)
	return re.MatchString(query)
}
