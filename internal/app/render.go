package main

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
