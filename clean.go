package main

import (
	"strings"
)

const (
	// allowedCharacters is a list of allowed characters outside of a quoted string, the first 6 being json structural characters.
	// https://datatracker.ietf.org/doc/html/rfc8259#section-2
	allowedCharacters = "[{]}:,\"1234567890.-"
)

func clean(input string) string {
	var insideQuotes bool
	var output string

	for _, c := range input {
		if c == '\\' {
			continue
		}

		if c == '"' {
			insideQuotes = !insideQuotes
		}

		if !insideQuotes && !strings.ContainsRune(allowedCharacters, c) {
			continue
		}

		output += string(c)
	}

	return output
}
