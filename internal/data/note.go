package data

import (
	"AhmadAbdelrazik/mark2right/internal/data/validator"
	"database/sql"
	"errors"
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

func (n *Note) Render(renderer IRender) string {
	return renderer.Render(n.Note)
}

func (n *Note) CheckSpelling(checker ILanguageChecker) []string {
	return checker.CheckSpelling(n.Note)
}

type NoteModel struct {
	DB *sql.DB
}

func (n *NoteModel) Insert(note *Note) error {
	query := `
	INSERT INTO notes (note, created_at)
	VALUES ($1, $2)
	RETURINING note_id`

	return n.DB.QueryRow(query, note.Note, note.CreatedAt).Scan(&note.NoteID)
}

func (n *NoteModel) Get(id int64) (*Note, error) {
	query := `
	SELECT note, created_at FROM notes
	WHERE id = $1`

	note := &Note{
		NoteID: id,
	}

	err := n.DB.QueryRow(query, id).Scan(&note.Note, &note.CreatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNoRecord
		default:
			return nil, err
		}
	}

	return note, nil
}

func (n *NoteModel) Update(note *Note) error {
	return nil
}

func (n *NoteModel) Delete(id int64) error {
	return nil
}
