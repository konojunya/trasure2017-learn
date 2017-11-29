package model

import "time"

type Report struct {
	ReportId int64      `json:"report_id"`
	Title    string     `json:"title"`
	Body     string     `json:"body"`
	Created  *time.Time `json:"created"`
	Updated  *time.Time `json:"updated"`
}
