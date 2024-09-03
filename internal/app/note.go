package main

import "time"

type ILanguageChecker interface {
	CheckSpelling(string) []string
}

type INote interface {
	Render() string
	CheckSpelling() []string
}

type Note struct {
	NoteID    int
	Note      string
	CreatedAt time.Time
	Renderer  IRender
	Checker   ILanguageChecker
}

type Checker struct {
	dictionaries [][]string
}
