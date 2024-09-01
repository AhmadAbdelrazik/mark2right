package api

import (
	"AhmadAbdelrazik/mark2right/internal/dictionary"

	"strings"
)

func CheckText(given string, d *dictionary.Dictionary) []string {
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

	regexp, _ := CompileRegex()

	for _, w := range words {
		if isAlphaNum(w) {
			if !d.Search(w) {
				badWords = append(badWords, w)
			}
		} else {
			if CheckRegex(regexp, w) {
				continue
			} else {
				badWords = append(badWords, w)
			}
		}
	}
	return badWords
}
