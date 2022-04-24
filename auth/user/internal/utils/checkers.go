package utils

import (
	"net/mail"
	"regexp"
)

var regexPhone = "^[0-9]{10,15}$"

func IsValidPhone(phones ...string) bool {
	for _, phone := range phones {
		match, _ := regexp.MatchString(regexPhone, phone)
		if !match {
			return false
		}
	}

	return true
}

func IsValidEmail(emails ...string) bool {
	for _, email := range emails {
		if _, err := mail.ParseAddress(email); err != nil {
			return false
		}
	}

	return true
}
