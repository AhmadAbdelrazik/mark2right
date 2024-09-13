package app

import (
	"fmt"
	"regexp"
)

type HeaderRenderer struct {
	headerRegex    *regexp.Regexp
	startSkipRegex *regexp.Regexp
	endSkipRegex   *regexp.Regexp
}

func NewHeaderRenderer() *HeaderRenderer {
	header := &HeaderRenderer{}

	header.headerRegex = regexp.MustCompile(`(^|\n)#{1,6} .*`)

	return header
}

func (r *HeaderRenderer) Render(input string) string {
	output := input

	for {
		loc := r.headerRegex.FindStringIndex(output)
		if loc == nil {
			break
		}
		begin, end := loc[0], loc[1]
		if output[begin] != '#' {
			begin++
		}

		i := 0
		for output[begin+i] == '#' {
			i++
		}

		renderedLine := fmt.Sprintf("<h%d>%s</h%d>", i, output[begin+i+1:end], i)
		output = output[:begin] + renderedLine + output[end:]

	}

	return output
}
