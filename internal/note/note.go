package note

import (
	"AhmadAbdelrazik/mark2right/internal/note/renderer"
	spellingchecker "AhmadAbdelrazik/mark2right/internal/note/spellingChecker"
	"AhmadAbdelrazik/mark2right/internal/note/validator"
	"time"

	_ "github.com/lib/pq"
)

type Note struct {
	NoteID    int64     `json:"note_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	Version   int64     `json:"version"`
}

func ValidateNote(v *validator.Validator, note *Note) {
	v.Check(note.Note != "", "note", "note can't be empty")
	v.Check(len(note.Note) <= 10_000, "note", "maximum length of note is 10,000 bytes")
}

func RenderNote(note Note) string {
	return renderer.Render(note.Note)
}

func CheckNoteSpelling(note Note) []string {
	return spellingchecker.CheckSpelling(note.Note)
}
