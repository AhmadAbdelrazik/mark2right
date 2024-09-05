package app

import (
	_ "embed"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed words.txt
var words string

//go:embed names.txt
var names string

type ILanguageChecker interface {
	CheckSpelling(string) []string
}

type Checker struct {
	words []string
	names []string
}

func (c *Checker) CheckSpelling(input string) []string {
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

		if !c.search(w) {
			badWords = append(badWords, w)
		}
	}
	return badWords
}

func NewChecker() ILanguageChecker {
	checker := &Checker{}
	checker.words = strings.Split(strings.ToLower(words), "\n")
	checker.names = strings.Split(strings.ToLower(names), "\n")

	sort.Slice(checker.words, func(i, j int) bool {
		return checker.words[i] < checker.words[j]
	})

	sort.Slice(checker.names, func(i, j int) bool {
		return checker.words[i] < checker.words[j]
	})

	return checker
}

func (c *Checker) search(word string) bool {
	word = strings.ToLower(word)
	// Check if it's a number
	if _, err := strconv.Atoi(word); err == nil {
		return true
	}

	// check if it's a name
	if isName, _ := binarySearch(word, c.names); isName {
		return true
	}

	// i : 182744, is : 198014
	if isWord, _ := binarySearch(word, c.words); isWord {
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
