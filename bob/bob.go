// Package bob replies with some predefined phrases depend on the text received
package bob

import (
	"strings"
	"unicode"
)

// Hey replies with some predefined texts depend on the remark given
func Hey(remark string) string {
	//Clean string with empty spaces
	remark = strings.Trim(remark, " ")

	//Check the case when all are spaces or tabs
	isEmpty := strings.Join(strings.Fields(remark), "")
	if len(isEmpty) == 0 {
		return "Fine. Be that way!"
	}

	//Check the case when all letters are uppercase
	var isYelling bool
	for _, r := range remark {
		if unicode.IsLower(r) {
			isYelling = false
			break
		}
		if unicode.IsUpper(r) {
			isYelling = true
		}
	}
	if isYelling == true {
		return "Whoa, chill out!"
	}

	//Check the case when is a question
	lastCharacter := remark[len(remark)-1:]
	if lastCharacter == "?" {
		return "Sure."
	}

	return "Whatever."
}
