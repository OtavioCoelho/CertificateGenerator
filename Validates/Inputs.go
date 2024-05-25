package Validates

import (
	"net/url"
	"strconv"
)

func InputArrayString(input string) []string {
	if input == "" {
		return nil
	}
	return []string{input}
}

func InputInt(yearsValidate string) int {
	if yearsValidate == "" {
		return 1
	}
	yers, err := strconv.Atoi(yearsValidate)
	if err != nil || yers <= 0 {
		return 1
	}
	return yers
}

func ValidateInputArrayUrl(urlString string) []*url.URL {
	if urlString == "" {
		return nil
	}
	parsedURI, err := url.Parse(urlString)
	if err != nil {
		return nil
	}
	return []*url.URL{parsedURI}
}
