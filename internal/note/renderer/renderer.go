package renderer

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	// Code Block Regex
	multiLineOpenRegex  = regexp.MustCompile("^```")
	multiLineCloseRegex = regexp.MustCompile("^```$")

	// Inline Code regex
	inlineOneTickRegex = regexp.MustCompile(`(^|[^\\\x60])\x60[^\x60]*\x60`)
	inlineTwoTickRegex = regexp.MustCompile(`(^|[^\\\x60])\x60{2}[^\x60]*\x60{2}`)

	linkRegex   = regexp.MustCompile(`(^|[^\\])\[[^\n\]]*\]\([^\)\n]*\)`)
	headerRegex = regexp.MustCompile(`(^|\n)#{1,6} .*`)

	// list regex
	unorderedRegex = regexp.MustCompile(`^(  ){0,5}- `)
	orderedRegex   = regexp.MustCompile(`^(  ){0,5}\d*\. `)

	// font regex
	italicRegex  = regexp.MustCompile(`\*([^<* ]|<[^* i]|<i[^* >])[^\n*]*[^ *\n]\*`)
	italicRegex2 = regexp.MustCompile(`\*[^ *\n]\*`)
	boldRegex    = regexp.MustCompile(`(^|[^\\])\*{2}([^< *\n]|<[^ b*\n]|<b[^ >*\n])([^\n*]|\\\*)*[^ *\n\\]\*{2}`)
	boldRegex2   = regexp.MustCompile(`\*{2}[^ *\n]\*{2}`)
)

func renderers() []func(string) string {
	var renders []func(string) string

	renders = append(renders, renderInlineCode)
	renders = append(renders, renderLink)
	renders = append(renders, renderList)
	renders = append(renders, renderHeader)
	renders = append(renders, renderFont)

	return renders
}

func Render(input string) string {
	var outputs []string
	multiLineActive := false

	renderers := renderers()

	for _, line := range strings.Split(input, "\n") {
		// If the multi line closing pattern appears, close the
		// multiline code block.
		if multiLineCloseRegex.MatchString(line) && multiLineActive {
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
		if multiLineOpenRegex.MatchString(line) {
			multiLineActive = true

			outputs = append(outputs, "<code>")
			// The remaining line can be used for syntax
			// highlighting functionality later on:
			// outputs = append(outputs, line[3:])
			continue
		}

		for _, rend := range renderers {
			line = rend(line)
		}

		outputs = append(outputs, line)
	}

	return strings.Join(outputs, "\n")

}

func renderInlineCode(input string) string {
	var output string

	// Looping on the line until there is no inline code blocks
	for {
		loc := inlineTwoTickRegex.FindStringIndex(input)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		// Used to check if there was a backslash behind the tick mark.
		if input[begin] != '`' {
			begin++
		}

		input = input[:begin] + "<code>" + input[begin+2:end-2] + "</code>" + input[end:]
	}

	// Looping on the line until there is no inline code blocks
	for {
		loc := inlineOneTickRegex.FindStringIndex(input)
		if loc == nil {
			break
		}

		// Used to check if there was a backslash behind the tick mark.
		begin, end := loc[0], loc[1]
		if input[begin] != '`' {
			begin++
		}

		input = input[:begin] + "<code>" + input[begin+1:end-1] + "</code>" + input[end:]
	}

	output += input

	return output
}

func renderLink(input string) string {
	output := input

	for {
		loc := linkRegex.FindStringIndex(output)
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

func renderList(input string) string {
	var output string

	calculateListLevel := func(input string) int {
		trimmedInput := strings.TrimLeft(input, " ")
		spaces := len(input) - len(trimmedInput)
		return (spaces / 2) + 1
	}

	cleanseLine := func(input string) string {
		input = strings.TrimSpace(input)
		words := strings.Split(input, " ")
		return strings.Join(words[1:], " ")
	}

	// Divide the input to lines.
	for _, line := range strings.Split(input, "\n") {
		// check for a list pattern.
		if loc := orderedRegex.FindStringIndex(line); loc != nil {
			// calculate the line level.
			level := calculateListLevel(line)
			for range level - 1 {
				output += "<ul>\n"
			}

			line = strings.TrimSpace(line)

			// Extract the number
			number := strings.Split(line, ".")[0]

			output += fmt.Sprintf("<ol start=%q>", number)
			output += "<li>" + cleanseLine(line) + "</li></ol>\n"

			for range level - 1 {
				output += "</ul>\n"
			}
			continue
		}

		if loc := unorderedRegex.FindStringIndex(line); loc != nil {
			// calculate the line level.
			level := calculateListLevel(line)
			for range level - 1 {
				output += "<ul>\n"
			}

			line = strings.TrimSpace(line)

			output += "<ul><li>" + cleanseLine(line) + "</li></ul>\n"

			for range level - 1 {
				output += "</ul>\n"
			}
			continue
		}

		output += line
	}

	return strings.TrimSuffix(output, "\n")
}

func renderHeader(input string) string {
	output := input

	for {
		loc := headerRegex.FindStringIndex(output)
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

func renderFont(input string) string {
	output := input

	for {
		loc := boldRegex.FindStringIndex(output)
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
		loc := boldRegex2.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		output = output[:begin] + "<b>" + output[begin+2:end-2] + "</b>" + output[end:]
	}

	for {
		loc := italicRegex.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		output = output[:begin] + "<i>" + output[begin+1:end-1] + "</i>" + output[end:]
	}

	for {
		loc := italicRegex2.FindStringIndex(output)
		if loc == nil {
			break
		}

		begin, end := loc[0], loc[1]
		output = output[:begin] + "<i>" + output[begin+1:end-1] + "</i>" + output[end:]
	}

	return output
}
