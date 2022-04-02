package utils

import (
	"regexp"
	"strings"
)

const IpDigit = "(\\d{1,2}|1\\d{1,2}|2[0-4]\\d|25[0-5])"

type IPChecker struct {
	Whitelist []string
}

func (i IPChecker) IsAllowed(ip string) bool {
	if strings.HasPrefix(ip, "[::1]") || strings.HasPrefix(ip, "127.0.0.1") {
		return true
	}
	ip = strings.Split(ip, ":")[0]

	for _, allowedIp := range i.Whitelist {
		pattern := strings.ReplaceAll(allowedIp, ".", "\\.")
		pattern = strings.ReplaceAll(pattern, "*", IpDigit)
		pattern = "^" + pattern + "$"

		isValid, _ := regexp.MatchString(pattern, ip)
		if isValid {
			return true
		}
	}

	return false
}
