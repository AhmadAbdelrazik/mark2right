package app

import (
	"fmt"
	"regexp"
	"strings"
)

type HeaderRenderer struct {
	Regex *regexp.Regexp
}

func NewHeaderRenderer() (IRender, error) {
	header := &HeaderRenderer{}

	regex, err := regexp.Compile("^#{1,6} .*")
	if err != nil {
		return nil, err
	}
	header.Regex = regex

	return header, nil
}

func (r *HeaderRenderer) Render(input string) string {
	if !r.Regex.MatchString(input) {
		return input
	}

	inputs := strings.Split(input, " ")
	hashes := len(inputs[0])

	output := fmt.Sprintf("<h%d>%s</h%d>", hashes, strings.Join(inputs[1:], " "), hashes)

	return output
}

type FontRenderer struct {
	Regex *regexp.Regexp
}

func NewFontRenderer() (IRender, error) {
	font := &FontRenderer{}

	regex, err := regexp.Compile("")
	if err != nil {
		return nil, err
	}
	font.Regex = regex

	return font, nil
}

func (r *FontRenderer) Render(input string) string {
	return ""
}
