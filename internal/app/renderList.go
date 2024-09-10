package app

import (
	"fmt"
	"regexp"
	"strings"
)

type ListRenderer struct {
	orderedRegex   *regexp.Regexp
	unorderedRegex *regexp.Regexp
}

func NewListRenderer() (*ListRenderer, error) {
	list := &ListRenderer{}

	unorderedRegex, err := regexp.Compile(`^(  ){0,5}- `)
	if err != nil {
		return nil, err
	}

	list.unorderedRegex = unorderedRegex

	orderedRegex, err := regexp.Compile(`^(  ){0,5}\d*\. `)
	if err != nil {
		return nil, err
	}

	list.orderedRegex = orderedRegex

	return list, nil
}

func (r *ListRenderer) Render(input string) string {
	var output string

	// Divide the input to lines.
	for _, line := range strings.Split(input, "\n") {
		// check for a list pattern.
		if loc := r.orderedRegex.FindStringIndex(line); loc != nil {
			// calculate the line level.
			level := r.CalculateListLevel(line)
			for range level - 1 {
				output += "<ul>\n"
			}

			line = strings.TrimSpace(line)

			// Extract the number
			number := strings.Split(line, ".")[0]

			output += fmt.Sprintf("<ol start=%q>", number)
			output += "<li>" + r.CleanseLine(line) + "</li></ol>\n"

			for range level - 1 {
				output += "</ul>\n"
			}
		}

		if loc := r.unorderedRegex.FindStringIndex(line); loc != nil {
			// calculate the line level.
			level := r.CalculateListLevel(line)
			for range level - 1 {
				output += "<ul>\n"
			}

			line = strings.TrimSpace(line)

			output += "<ul><li>" + r.CleanseLine(line) + "</li></ul>\n"

			for range level - 1 {
				output += "</ul>\n"
			}
		}
	}

	return output
}

func (r *ListRenderer) CalculateListLevel(input string) int {
	trimmedInput := strings.TrimLeft(input, " ")
	spaces := len(input) - len(trimmedInput)
	return (spaces / 2) + 1
}

// CleanseLine Separate the list mark "1. " or "- " from the line and return
// the line only
func (r *ListRenderer) CleanseLine(input string) string {
	input = strings.TrimSpace(input)
	words := strings.Split(input, " ")
	return strings.Join(words[1:], " ")
}
