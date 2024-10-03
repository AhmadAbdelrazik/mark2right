package note

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRecord     = errors.New("record not found")
	ErrEditConflict = errors.New("edit confilct")
)

type Models struct {
	Notes NoteModel
}

func NewModels(db *sql.DB) *Models {
	return &Models{
		Notes: NoteModel{DB: db},
	}
}
