package app

import "regexp"

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
