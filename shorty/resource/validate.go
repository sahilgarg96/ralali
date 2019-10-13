package resource

import (
	"net/url"
	"regexp"
	"strings"
)

const (
	URL string = `^[0-9a-zA-Z_]{6}$`
)

func ValidateURL(value string) bool {

	rgxURL := regexp.MustCompile(URL)

	if value == "" || strings.HasPrefix(value, ".") {
		return false
	}
	tempValue := value
	// Validate URLs that do not start with a scheme
	if strings.Contains(value, ":") && !strings.Contains(value, "://") {
		tempValue = "http://" + value
	}
	u, err := url.Parse(tempValue)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	return rgxURL.MatchString(value)
}
