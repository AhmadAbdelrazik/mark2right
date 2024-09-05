package app

import "time"

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
