package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snip struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnipModel struct {
	DB *sql.DB
}

func (m *SnipModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snips (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnipModel) Get(id int) (Snip, error) {
	stmt := `SELECT id, title, content, created, expires FROM snips WHERE expires > UTC_TIMESTAMP() AND id = ?`
	row := m.DB.QueryRow(stmt, id)
	var s Snip
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snip{}, ErrNoRecord
		} else {
			return Snip{}, err
		}
	}
	return s, nil
}

func (m *SnipModel) Latest() ([]Snip, error) {
	stmt := `SELECT id, title, content, created, expires FROM snips WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var snips []Snip
	for rows.Next() {
		var s Snip
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snips = append(snips, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snips, nil
}
