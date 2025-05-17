package rules

import (
	"regexp"
)

func GetRegexes() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?i)<script[^>]*>`),  // XSS
		regexp.MustCompile(`(?i)union\s+select`), // SQLi
		regexp.MustCompile(`(?i)onerror\s*=`),    // XSS
		regexp.MustCompile(`(?i)eval\((.*?)\)`),  // dangerous eval
	}
}
