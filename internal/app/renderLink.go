package app

import (
	"regexp"
)

type LinkRenderer struct {
	regex *regexp.Regexp
}

func NewLinkRenderer() *LinkRenderer {
	r := &LinkRenderer{}

	r.regex = regexp.MustCompile(`(^|[^\\])\[[^\n\]]*\]\([^\)\n]*\)`)

	return r
}

func (r *LinkRenderer) Render(input string) string {
	output := input

	for {
		loc := r.regex.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		if output[begin] != '[' {
			begin++
		}

		var text, link string

		i := begin
		for {

			if output[i] == ']' {
				text = output[begin+1 : i]
				link = output[i+2 : end-1]
				break
			}

			i++
		}

		output = output[:begin] + `<a href="` + link + `">` + text + "</a>" + output[end:]

	}

	return output
}
