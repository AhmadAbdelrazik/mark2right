package app

import (
	"regexp"
	"strings"
)

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
