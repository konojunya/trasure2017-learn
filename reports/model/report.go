package model

import (
	"database/sql"
)

func (r *Report) Insert(tx *sql.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	insert into daily_report (title, body)
	values(?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(r.Title, r.Body)
}
