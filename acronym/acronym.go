// Package acronym help doing operations with acronyms
package acronym

import (
  "strings"
)

// Abbreviate converts a phrase into acronym
func Abbreviate(s string) string {
	var result string
  s = strings.Replace(s, "-", " ", -1)
	words := strings.Split(s, " ")

	for _ , w := range words {
		result += strings.ToUpper(string(w[0]))
	}

	return result
}
