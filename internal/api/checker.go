package api

import (
	"strings"
)

func (a *Application) CheckText(given string) []string {
	// Get each word out
	words := strings.Split(given, " ")

	// trim the signs
	trimmers := "'.,!?@#$%^&*/*+`\""
	for i := range words {
		for _, t := range trimmers {
			words[i] = strings.Trim(words[i], string(t))
			words[i] = strings.TrimSpace(words[i])
		}
	}

	badWords := []string{}

	for _, w := range words {
		if isAlphaNum(w) {
			if !a.Dictionary.Search(w) {
				badWords = append(badWords, w)
			}
		} else {
			if CheckRegex(a.Regex, w) {
				continue
			} else {
				badWords = append(badWords, w)
			}
		}
	}
	return badWords
}
