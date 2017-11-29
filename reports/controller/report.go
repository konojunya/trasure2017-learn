package controller

import (
	"database/sql"
	"flag"

	"github.com/konojunya/voyage-group/reports/model"
)

type Report struct {
	DB *sql.DB
}

func (r *Report) New(m *model.Report) error {
	if err := TXHandler(r.DB, func(tx *sql.Tx) error {
		_, err := m.Insert(tx)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *Report) Save() error {

	var (
		title string
		body  string
	)

	// title
	flag.StringVar(&title, "t", "日報", "daily report")

	// body
	flag.StringVar(&body, "b", "", "report body")
	flag.Parse()

	var report model.Report
	report.Title = "title"
	report.Body = "body"

	return r.New(&report)
}
