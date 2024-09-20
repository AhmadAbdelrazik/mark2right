package data

import (
	"fmt"
	"regexp"
	"strings"
)

type IRender interface {
	Render(string) string
}

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

type CodeRenderer struct {
	inlineOneTickRegex  *regexp.Regexp
	inlineTwoTickRegex  *regexp.Regexp
	multiLineOpenRegex  *regexp.Regexp
	multiLineCloseRegex *regexp.Regexp
}

func NewCodeRenderer() *CodeRenderer {
	r := &CodeRenderer{}

	r.inlineOneTickRegex = regexp.MustCompile(`(^|[^\\\x60])\x60[^\x60]*\x60`)
	r.inlineTwoTickRegex = regexp.MustCompile(`(^|[^\\\x60])\x60{2}[^\x60]*\x60{2}`)
	r.multiLineOpenRegex = regexp.MustCompile("^```")
	r.multiLineCloseRegex = regexp.MustCompile("^```$")

	return r
}

func (r *CodeRenderer) Render(input string) string {
	var output string
	multiLineActive := false

	for _, line := range strings.Split(input, "\n") {
		// If the multi line closing pattern appears, close the
		// multiline code block.
		if r.multiLineCloseRegex.MatchString(line) && multiLineActive {
			multiLineActive = false
			output += "</code>"
			continue
		}

		// while multiline active, no styling is done inside.
		if multiLineActive {
			output += line + "\n"
			continue
		}

		// if multiline opening is found, start a multi line code block
		if r.multiLineOpenRegex.MatchString(line) {
			multiLineActive = true

			output += "<code>\n"
			output += line[3:]
			continue
		}

		// Looping on the line until there is no inline code blocks
		for {
			loc := r.inlineTwoTickRegex.FindStringIndex(line)
			if loc == nil {
				break
			}

			begin, end := loc[0], loc[1]
			// Used to check if there was a backslash behind the tick mark.
			if line[begin] != '`' {
				begin++
			}

			line = line[:begin] + "<code>" + line[begin+2:end-2] + "</code>" + line[end:]
		}

		// Looping on the line until there is no inline code blocks
		for {
			loc := r.inlineOneTickRegex.FindStringIndex(line)
			if loc == nil {
				break
			}

			// Used to check if there was a backslash behind the tick mark.
			begin, end := loc[0], loc[1]
			if line[begin] != '`' {
				begin++
			}

			line = line[:begin] + "<code>" + line[begin+1:end-1] + "</code>" + line[end:]
		}

		output += line
	}

	return output
}

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

type ListRenderer struct {
	orderedRegex   *regexp.Regexp
	unorderedRegex *regexp.Regexp
}

func NewListRenderer() *ListRenderer {
	list := &ListRenderer{}

	list.unorderedRegex = regexp.MustCompile(`^(  ){0,5}- `)
	list.orderedRegex = regexp.MustCompile(`^(  ){0,5}\d*\. `)

	return list
}

func (r *ListRenderer) Render(input string) string {
	var output string

	// Divide the input to lines.
	for _, line := range strings.Split(input, "\n") {
		// check for a list pattern.
		if loc := r.orderedRegex.FindStringIndex(line); loc != nil {
			// calculate the line level.
			level := r.CalculateListLevel(line)
			for range level - 1 {
				output += "<ul>\n"
			}

			line = strings.TrimSpace(line)

			// Extract the number
			number := strings.Split(line, ".")[0]

			output += fmt.Sprintf("<ol start=%q>", number)
			output += "<li>" + r.CleanseLine(line) + "</li></ol>\n"

			for range level - 1 {
				output += "</ul>\n"
			}
			continue
		}

		if loc := r.unorderedRegex.FindStringIndex(line); loc != nil {
			// calculate the line level.
			level := r.CalculateListLevel(line)
			for range level - 1 {
				output += "<ul>\n"
			}

			line = strings.TrimSpace(line)

			output += "<ul><li>" + r.CleanseLine(line) + "</li></ul>\n"

			for range level - 1 {
				output += "</ul>\n"
			}
			continue
		}

		output += line
	}

	return output
}

func (r *ListRenderer) CalculateListLevel(input string) int {
	trimmedInput := strings.TrimLeft(input, " ")
	spaces := len(input) - len(trimmedInput)
	return (spaces / 2) + 1
}

// CleanseLine Separate the list mark "1. " or "- " from the line and return
// the line only
func (r *ListRenderer) CleanseLine(input string) string {
	input = strings.TrimSpace(input)
	words := strings.Split(input, " ")
	return strings.Join(words[1:], " ")
}

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

type FontRenderer struct {
	boldRegex    *regexp.Regexp
	boldRegex2   *regexp.Regexp
	italicRegex  *regexp.Regexp
	italicRegex2 *regexp.Regexp
}

func NewFontRenderer() *FontRenderer {
	font := &FontRenderer{}

	font.italicRegex = regexp.MustCompile(`\*([^<* ]|<[^* i]|<i[^* >])[^\n*]*[^ *\n]\*`)
	font.italicRegex2 = regexp.MustCompile(`\*[^ *\n]\*`)
	font.boldRegex = regexp.MustCompile(`(^|[^\\])\*{2}([^< *\n]|<[^ b*\n]|<b[^ >*\n])([^\n*]|\\\*)*[^ *\n\\]\*{2}`)
	font.boldRegex2 = regexp.MustCompile(`\*{2}[^ *\n]\*{2}`)

	return font
}

func (r *FontRenderer) Render(input string) string {
	output := input

	for {
		loc := r.boldRegex.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		if output[begin] != '*' || output[begin+2] == '*' {
			begin++
		}
		output = output[:begin] + "<b>" + output[begin+2:end-2] + "</b>" + output[end:]
	}

	for {
		loc := r.boldRegex2.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		output = output[:begin] + "<b>" + output[begin+2:end-2] + "</b>" + output[end:]
	}

	for {
		loc := r.italicRegex.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		output = output[:begin] + "<i>" + output[begin+1:end-1] + "</i>" + output[end:]
	}

	for {
		loc := r.italicRegex2.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		output = output[:begin] + "<i>" + output[begin+1:end-1] + "</i>" + output[end:]
	}

	return output
}
