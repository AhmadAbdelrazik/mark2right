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

func NewListRenderer() (IRender, error) {
	list := &ListRenderer{}

	unorderedRegex, err := regexp.Compile(`^(  ){0,5}- `)
	if err != nil {
		return nil, err
	}

	list.unorderedRegex = unorderedRegex

	orderedRegex, err := regexp.Compile(`^(  ){0,5}\d\. `)
	if err != nil {
		return nil, err
	}

	list.orderedRegex = orderedRegex

	return list, nil
}

func (r *ListRenderer) Render(input string) string {
	var output string
	var s Stack

	// Divide the input to lines.
	for _, line := range strings.Split(input, "\n") {

		// check for a list pattern.

		// ordered list check
		if loc := r.orderedRegex.FindStringIndex(line); loc != nil {
			// calculate the line level.
			level := r.CalculateListLevel(line)
			if !s.Empty() {
				output += r.Resolve(s, fmt.Sprintf("ol%d", level))
			}

			output += "<ol>\n"
			output += "<li>" + line + "</li>\n"

			s.Push(fmt.Sprintf("ol%d", level))
		}
	}

	return output
}

func (r *ListRenderer) CalculateListLevel(input string) int {
	trimmedInput := strings.TrimLeft(input, " ")
	spaces := len(input) - len(trimmedInput)
	return (spaces / 2) + 1
}

func (r *ListRenderer) Resolve(s Stack, level string) string {
	var output string

	for !s.Empty() {
		top, _ := s.Top()
		s.Pop()

		if top[:2] == level[:2] && top[2] <= level[2] {
			break
		}

		output += "</" + top[:2] + ">\n"
	}

	return output
}
