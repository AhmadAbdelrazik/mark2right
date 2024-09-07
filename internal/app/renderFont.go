package app

import "regexp"

type FontRenderer struct {
	boldRegex    *regexp.Regexp
	boldRegex2   *regexp.Regexp
	italicRegex  *regexp.Regexp
	italicRegex2 *regexp.Regexp
}

func NewFontRenderer() (IRender, error) {
	font := &FontRenderer{}

	italicRegex, err := regexp.Compile(`\*([^<* ]|<[^* i]|<i[^* >])[^\n*]*[^ *\n]\*`)
	if err != nil {
		return nil, err
	}
	font.italicRegex = italicRegex

	italicRegex2, err := regexp.Compile(`\*[^ *\n]\*`)
	if err != nil {
		return nil, err
	}
	font.italicRegex2 = italicRegex2

	boldRegex, err := regexp.Compile(`(^|[^\\])\*{2}([^< *\n]|<[^ b*\n]|<b[^ >*\n])([^\n*]|\\\*)*[^ *\n\\]\*{2}`)
	if err != nil {
		return nil, err
	}
	font.boldRegex = boldRegex

	boldRegex2, err := regexp.Compile(`\*{2}[^ *\n]\*{2}`)
	if err != nil {
		return nil, err
	}
	font.boldRegex2 = boldRegex2

	return font, nil
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
