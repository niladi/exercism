// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "regexp"

const (
	msgQuestion     = "Sure."
	msgYell         = "Whoa, chill out!"
	msgYellQuestion = "Calm down, I know what I'm doing!"
	msgNone         = "Fine. Be that way!"
	msgAnything     = "Whatever."
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	if len(remark) == 0 {
		return msgNone
	} else if i, _ := regexp.MatchString("^\\s+$", remark); i {
		return msgNone
	}
	// if matched, _ := regexp.MatchString("^([\\w\\s\\.\"',\\!\\?\\.])+$", remark); matched {
	quest, _ := regexp.MatchString("\\?\\s*$", remark)
	yell, _ := regexp.MatchString("^([^a-z]*[A-Z][^a-z]*)+$", remark)
	if quest && yell {
		return msgYellQuestion
	} else if yell {
		return msgYell
	} else if quest {
		return msgQuestion
	}

	return msgAnything
}
