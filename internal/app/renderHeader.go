package app

import (
	"fmt"
	"regexp"
	"strings"
)

type HeaderRenderer struct {
	regex *regexp.Regexp
}

func NewHeaderRenderer() (*HeaderRenderer, error) {
	header := &HeaderRenderer{}

	regex, err := regexp.Compile("^#{1,6} .*")
	if err != nil {
		return nil, err
	}
	header.regex = regex

	return header, nil
}

func (r *HeaderRenderer) Render(input string) string {
	var renderedLines []string
	for _, line := range strings.Split(input, "\n") {
		if !r.regex.MatchString(line) {
			renderedLines = append(renderedLines, line)
			continue
		}

		inputs := strings.Split(line, " ")
		hashes := len(inputs[0])

		renderedLine := fmt.Sprintf("<h%d>%s</h%d>", hashes, strings.Join(inputs[1:], " "), hashes)

		renderedLines = append(renderedLines, renderedLine)
	}

	output := strings.Join(renderedLines, "\n")

	return output
}
