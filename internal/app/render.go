package main

import (
	"strings"
)

// Rendering interface, it takes the string that should be rendered and
// returns a rendered string
type IRender interface {
	Render(string) string
}

type Renderer struct {
	Rules []*RenderRule
}

func (r *Renderer) Render(input string) string {
	output := ""
	for _, line := range strings.Split(input, "\n") {
		for _, rule := range r.Rules {
			line = rule.Render(line)
		}
		output += line
	}

	return output
}

func NewRenderer() IRender {
	renderer := &Renderer{}

	return renderer
}
