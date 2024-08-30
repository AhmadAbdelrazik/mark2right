package dictionary

import (
	"os"
	"strings"
)

type Dictionary []string

func NewDictionary() (Dictionary, error) {
	dictRaw, err := os.ReadFile("words.txt")
	if err != nil {
		return nil, err
	}

	dictList := strings.Split(string(dictRaw), "\n")
	dictionary := make(Dictionary, len(dictList))

	for i, w := range dictList {
		dictionary[i] = w
	}

	return dictionary, nil
}

func (d Dictionary) Search(word string) bool {
	word = strings.ToLower(word)
	l, r := 0, len(d)-1

	for l <= r {
		m := (l + r) / 2
		if d[m] > word {
			r = m - 1
		} else if d[m] < word {
			l = m + 1
		} else if d[m] == word {
			return true
		}
	}
	return false
}
