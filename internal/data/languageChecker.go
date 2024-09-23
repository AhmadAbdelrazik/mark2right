package data

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed words.txt
var dictionaryWords string

//go:embed names.txt
var dictionaryNames string

var (
	allWords = strings.Split(dictionaryWords, "\n")
	allNames = strings.Split(dictionaryNames, "\n")
)

func CheckSpelling(input string) []string {
	// Get each word out
	words := strings.Split(input, " ")

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
		if !isAlphaNum(w) {
			continue
		}

		if !search(w) {
			badWords = append(badWords, w)
		}
	}
	return badWords
}

func search(word string) bool {
	word = strings.ToLower(word)
	// Check if it's a number
	if _, err := strconv.Atoi(word); err == nil {
		return true
	}

	// check if it's a name
	if isName, _ := binarySearch(word, allNames); isName {
		return true
	}

	// i : 182744, is : 198014
	if isWord, _ := binarySearch(word, allWords); isWord {
		return true
	}

	return false
}

func binarySearch(target string, list []string) (bool, int) {
	l, r := 0, len(list)-1
	target = strings.ToLower(target)

	for l <= r {
		m := (l + r) / 2
		word := strings.ToLower(list[m])
		if word > target {
			r = m - 1
		} else if word < target {
			l = m + 1
		} else if word == target {
			return true, m
		}
	}
	return false, l
}

// Need to be fixed. We mustn't compile the regex each time.
func isAlphaNum(w string) bool {
	if r, err := regexp.Compile("^[a-zA-Z0-9_]*$"); err != nil {
		panic(err)
	} else {
		return r.MatchString(w)
	}
}
