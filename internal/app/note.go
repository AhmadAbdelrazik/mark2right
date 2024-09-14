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
	Renderers IRender
	Checker   ILanguageChecker
}

func NewNote(noteText string, renderer IRender, checker ILanguageChecker) *Note {
	note := &Note{
		Note:      noteText,
		CreatedAt: time.Now(),
		Renderers: renderer,
		Checker:   checker,
	}

	return note
}

func (n *Note) Render() string {
	return n.Renderers.Render(n.Note)
}

func (n *Note) CheckSpelling() []string {
	return n.Checker.CheckSpelling(n.Note)
}
