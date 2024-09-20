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

	for i, c := range input {
		append := true
		if c == '\\' && i > 0 && input[i-1] != '\\' {
			append = false
		}

		if c == '"' && i > 1 && input[i-1] == '\\' && input[i-2] != '\\' {
			insideQuotes = !insideQuotes
		}

		if !insideQuotes && !strings.ContainsRune(allowedCharacters, c) {
			append = false
		}

		if append {
			output += string(c)
		}
	}

	return output
}
