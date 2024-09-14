package app

import (
	"regexp"
	"strings"
)

type Renderer struct {
	renderers           []IRender
	multiLineOpenRegex  *regexp.Regexp
	multiLineCloseRegex *regexp.Regexp
}

func NewRenderer() *Renderer {
	r := &Renderer{}

	hr := NewHeaderRenderer()
	fr := NewFontRenderer()
	lr := NewListRenderer()
	cr := NewCodeRenderer()
	llr := NewLinkRenderer()

	r.renderers = append(r.renderers, hr)
	r.renderers = append(r.renderers, fr)
	r.renderers = append(r.renderers, lr)
	r.renderers = append(r.renderers, cr)
	r.renderers = append(r.renderers, llr)

	r.multiLineOpenRegex = regexp.MustCompile("^```")
	r.multiLineCloseRegex = regexp.MustCompile("^```$")

	return r
}

func (r *Renderer) Render(input string) string {
	var outputs []string
	multiLineActive := false

	for _, line := range strings.Split(input, "\n") {
		// If the multi line closing pattern appears, close the
		// multiline code block.
		if r.multiLineCloseRegex.MatchString(line) && multiLineActive {
			multiLineActive = false
			outputs = append(outputs, "</code>")
			continue
		}

		// while multiline active, no styling is done inside.
		if multiLineActive {
			outputs = append(outputs, line)
			continue
		}

		// if multiline opening is found, start a multi line code block
		if r.multiLineOpenRegex.MatchString(line) {
			multiLineActive = true

			outputs = append(outputs, "<code>")
			// The remaining line can be used for syntax
			// highlighting functionality later on:
			// outputs = append(outputs, line[3:])
			continue
		}

		for _, rend := range r.renderers {
			line = rend.Render(line)
		}

		outputs = append(outputs, line)
	}

	return strings.Join(outputs, "\n")

}
