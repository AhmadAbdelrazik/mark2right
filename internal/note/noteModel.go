package note

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type NoteModel struct {
	DB *sql.DB
}

func (n *NoteModel) Insert(note *Note) error {
	query := `
	INSERT INTO notes (note, created_at)
	VALUES ($1, $2)
	RETURINING note_id`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return n.DB.QueryRowContext(ctx, query, note.Note, note.CreatedAt).Scan(&note.NoteID)
}

func (n *NoteModel) Get(id int64) (*Note, error) {
	query := `
	SELECT note, created_at FROM notes
	WHERE note_id = $1`

	note := &Note{
		NoteID: id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := n.DB.QueryRowContext(ctx, query, id).Scan(&note.Note, &note.CreatedAt)
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
	query := `
	UPDATE notes
	SET note = $1, version = version + 1
	WHERE note_id = $2 AND version = $3
	RETURNING version`
	args := []interface{}{note.Note, note.NoteID, note.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := n.DB.QueryRowContext(ctx, query, args...).Scan(&note.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (n *NoteModel) Delete(id int64) error {
	query := `
	DELETE FROM notes
	WHERE note_id = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := n.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNoRecord
	}
	return nil
}
