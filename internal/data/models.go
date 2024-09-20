package data

import (
	"database/sql"
	"errors"
)

var (
	ErrNoRecord = errors.New("record not found")
)

type Models struct {
	Notes interface {
		Insert(note *Note) error
		Get(id int64) (*Note, error)
		Update(note *Note) error
		Delete(id int64) error
	}
}

func NewModels(db *sql.DB) *Models {
	return &Models{
		Notes: &NoteModel{DB: db},
	}
}
