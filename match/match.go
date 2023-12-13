// match the logs lines coming from vault
// and classify them in the right categories
package match

import (
	"regexp"
)

// Match is a struct that contains the regex and the category
// to match the log line against
type Match struct {
	Regex    *regexp.Regexp
	Category string
}

// Matches is a list of Match
type Matches []Match

// MatchLogLine takes a log line and returns the category
// it matches
func (m Matches) MatchLogLine(logLine string) string {
	for _, match := range m {
		if match.Regex.MatchString(logLine) {
			return match.Category
		}
	}
	return ""
}

// NewMatches returns a Matches object from a list of
// regex and categories
func NewMatches(matches map[string]string) (Matches, error) {
	var m Matches
	for regex, category := range matches {
		r, err := regexp.Compile(regex)
		if err != nil {
			return nil, err
		}
		m = append(m, Match{Regex: r, Category: category})
	}
	return m, nil
}
