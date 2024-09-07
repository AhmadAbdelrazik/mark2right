package app

import "regexp"

type ListRenderer struct {
	orderedRegex   *regexp.Regexp
	unorderedRegex *regexp.Regexp
}

func NewListRenderer() (IRender, error) {
	list := &ListRenderer{}

	unorderedRegex, err := regexp.Compile(`^(  ){0,5}- .*$`)
	if err != nil {
		return nil, err
	}

	list.unorderedRegex = unorderedRegex

	orderedRegex, err := regexp.Compile(`^(  ){0,5}\d\. .*$`)
	if err != nil {
		return nil, err
	}

	list.orderedRegex = orderedRegex

	return list, nil
}

func (r *ListRenderer) Render(input string) string {
	output := input

	return output

}
