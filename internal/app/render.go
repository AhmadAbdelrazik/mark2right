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
	BoldRegex    *regexp.Regexp
	BoldRegex2   *regexp.Regexp
	ItalicRegex  *regexp.Regexp
	ItalicRegex2 *regexp.Regexp
}

func NewFontRenderer() (IRender, error) {
	font := &FontRenderer{}

	italicRegex, err := regexp.Compile(`\*([^<* ]|<[^* i]|<i[^* >])[^\n*]*[^ *\n]\*`)
	if err != nil {
		return nil, err
	}
	font.ItalicRegex = italicRegex

	italicRegex2, err := regexp.Compile(`\*[^ *\n]\*`)
	if err != nil {
		return nil, err
	}
	font.ItalicRegex2 = italicRegex2

	boldRegex, err := regexp.Compile(`\*{2}([^< *\n]|<[^ b*\n]|<b[^ >*\n])[^\n*]*[^ *\n]\*{2}`)
	if err != nil {
		return nil, err
	}
	font.BoldRegex = boldRegex

	boldRegex2, err := regexp.Compile(`\*{2}[^ *\n]\*{2}`)
	if err != nil {
		return nil, err
	}
	font.BoldRegex2 = boldRegex2

	return font, nil
}

func (r *FontRenderer) Render(input string) string {
	output := input

	for {
		loc := r.BoldRegex.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		output = output[:begin] + "<b>" + output[begin+2:end-2] + "</b>" + output[end:]
	}

	for {
		loc := r.BoldRegex2.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		output = output[:begin] + "<b>" + output[begin+2:end-2] + "</b>" + output[end:]
	}

	for {
		loc := r.ItalicRegex.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		output = output[:begin] + "<i>" + output[begin+1:end-1] + "</i>" + output[end:]
	}

	for {
		loc := r.ItalicRegex2.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		output = output[:begin] + "<i>" + output[begin+1:end-1] + "</i>" + output[end:]
	}

	return output
}
