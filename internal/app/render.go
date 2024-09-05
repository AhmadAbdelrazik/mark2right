package app

import (
	"fmt"
	"regexp"
	"strings"
)

// Rendering interface, it takes the string that should be rendered and
// returns a rendered string
type IRender interface {
	Render(string) string
}

// Template
// type Renderer struct {
// 	Regex *regexp.Regexp
// }
//
// func (r *Renderer) Render(input string) string {
// 	if !r.Regex.MatchString(input) {
// 		return input
// 	}
//
// 	// Implement the renderer here
//
// 	return ""
// }

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
