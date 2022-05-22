package utils

import (
	"regexp"
)

var (
	regexUrl = "^([a-zA-Z0-9|-]+\\.?){1,64}[[a-zA-Z0-9|-]+\\.[a-zA-Z]+$"
)

func IsValidUrl(urls ...string) bool {
	for _, url := range urls {
		match, _ := regexp.MatchString(regexUrl, url)
		if !match {
			return false
		}
	}

	return true
}
