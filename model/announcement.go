package model

import (
	"time"
)

//go:generate reform

// Announcement represents a row in announcement table.
//reform:announcement
type Announcement struct {
	ID       int32     `reform:"id,pk"`
	Date     time.Time `reform:"date"`
	Content  string    `reform:"content"`
	Markdown string    `reform:"markdown"`
}
