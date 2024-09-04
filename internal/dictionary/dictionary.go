package dictionary

import (
	"sort"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed words.txt
var words string

//go:embed names.txt
var names string

type Dictionary struct {
	words []string
	names []string
}

// NewDictionary Loads the dictionary slice to the program
func NewDictionary() (*Dictionary, error) {
	dict := &Dictionary{}
	dict.words = strings.Split(strings.ToLower(string(words)), "\n")
	dict.names = strings.Split(strings.ToLower(string(names)), "\n")

	sort.Slice(dict.words, func(i, j int) bool {
		return dict.words[i] < dict.words[j]
	})
	sort.Slice(dict.names, func(i, j int) bool {
		return dict.words[i] < dict.words[j]
	})
	return dict, nil
}

// Search search for a word in the dictionary. It's acceptable if it's a
// number, name or an english word
func (d *Dictionary) Search(word string) bool {
	word = strings.ToLower(word)
	// Check if it's a number
	if _, err := strconv.Atoi(word); err == nil {
		return true
	}

	// check if it's a name
	if isName, _ := binarySearch(word, d.names); isName {
		return true
	}

	// i : 182744, is : 198014
	if isWord, _ := binarySearch(word, d.words); isWord {
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
