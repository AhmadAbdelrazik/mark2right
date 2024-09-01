package api

import (
	"regexp"
)

func isAlphaNum(w string) bool {
	if r, err := regexp.Compile("^[a-zA-Z0-9_]*$"); err != nil {
		panic(err)
	} else {
		return r.MatchString(w)
	}
}

func CompileRegex() ([]*regexp.Regexp, error) {
	regexps := []*regexp.Regexp{}

	if emailRegex, err := regexp.Compile(`^([a-z0-9_\.\+-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$`); err != nil {
		return nil, err
	} else {
		regexps = append(regexps, emailRegex)
	}

	if urlRegex, err := regexp.Compile(`(https?:\/\/)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`); err != nil {
		return nil, err
	} else {
		regexps = append(regexps, urlRegex)
	}

	if dateRegex, err := regexp.Compile(`^\d{1,2}[-\/](Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)[-\/]\d*$`); err != nil {
		return nil, err
	} else {
		regexps = append(regexps, dateRegex)
	}

	if timeRegex, err := regexp.Compile(`(?:[01]\d|2[0123]):(?:[012345]\d):(?:[012345]\d)`); err != nil {
		return nil, err
	} else {
		regexps = append(regexps, timeRegex)
	}

	if slugRegex, err := regexp.Compile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`); err != nil {
		return nil, err
	} else {
		regexps = append(regexps, slugRegex)
	}

	if phoneRegex, err := regexp.Compile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`); err != nil {
		return nil, err
	} else {
		regexps = append(regexps, phoneRegex)
	}

	return regexps, nil
}

func CheckRegex(regexps []*regexp.Regexp, w string) bool {
	for _, r := range regexps {
		if r.MatchString(w) {
			return true
		}
	}

	return false
}
