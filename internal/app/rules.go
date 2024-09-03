package main

import "regexp"

type RenderRule struct {
	Regex    *regexp.Regexp
	Renderer func(string) string
}

func NewRenderRule(regexString string, renderer func(string) string) (*RenderRule, error) {

	regex, err := regexp.Compile(regexString)
	if err != nil {
		return nil, err
	}

	return &RenderRule{
		Regex:    regex,
		Renderer: renderer,
	}, nil
}

func (r *RenderRule) Render(input string) string {
	if !r.Regex.MatchString(input) {
		return input
	}

	return r.Renderer(input)
}
