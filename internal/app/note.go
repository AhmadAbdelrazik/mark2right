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
	Renderers []IRender
	Checker   ILanguageChecker
}

func NewNote(noteText string, renderers []IRender, checker ILanguageChecker) *Note {
	note := &Note{
		Note:      noteText,
		CreatedAt: time.Now(),
		Renderers: renderers,
		Checker:   checker,
	}

	return note
}

func (n *Note) Render() string {
	output := n.Note

	for _, r := range n.Renderers {
		output = r.Render(output)
	}

	return output
}

func (n *Note) CheckSpelling() []string {
	return n.Checker.CheckSpelling(n.Note)
}
