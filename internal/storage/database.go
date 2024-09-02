package storage

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrNoRecord = errors.New("no record")
)

type Note struct {
	NoteID    int
	Note      string
	CreatedAt time.Time
}

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(dsn string) (*MySQL, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	mysql := &MySQL{
		DB: db,
	}
	return mysql, nil
}

func (m *MySQL) CreateNote(note string) (int, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return 0, err
	}

	stmt := `INSERT INTO notes(note, created_at) VALUES (?,?)`

	result, err := tx.Exec(stmt, note, time.Now())
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(id), nil
}

func (m *MySQL) GetNotes() ([]Note, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	stmt := `SELECT note_id, note, created_at FROM notes`

	rows, err := tx.Query(stmt)
	if err != nil {
		return nil, err
	}

	notes := []Note{}
	for rows.Next() {
		note := Note{}
		err := rows.Scan(&note.NoteID, &note.Note, &note.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func (m *MySQL) GetNote(noteID int) (*Note, error) {
	tx, err := m.DB.Begin()
	if err != nil {
		return nil, err
	}

	stmt := `SELECT note, created_at FROM notes WHERE note_id = ?`

	row := tx.QueryRow(stmt, noteID)

	note := &Note{}
	err = row.Scan(&note.NoteID, &note.Note, &note.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return note, nil
}

func (m *MySQL) UpdateNote(note Note) error {
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	stmt := `UPDATE notes set note = ? WHERE note_id = ?`

	if _, err := tx.Exec(stmt, note); err != nil {
		tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoRecord
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (m *MySQL) DeleteNote(noteID int) error {
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}

	stmt := `DELETE FROM notes WHERE note_id = ?`

	if _, err := tx.Exec(stmt, noteID); err != nil {
		tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoRecord
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
