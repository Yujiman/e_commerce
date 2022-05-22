package utils

import (
	"encoding/json"
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

func IsValidJson(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}
