// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	re := regexp.MustCompile("([^A-Z'][A-Z]|^[A-Z])")
	letters := re.FindAllString(strings.ToUpper(s), -1)
	var b strings.Builder
	for _, v := range letters {
		b.WriteByte(v[len(v)-1])
	}
	return b.String()
}
