package app

import "time"

// Rendering interface, it takes the string that should be rendered and
// returns a rendered string
type IRender interface {
	Render(string) string
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
