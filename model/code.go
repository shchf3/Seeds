package model

import (
	"time"
)

//go:generate reform

// Code represents a row in code table.
//reform:code
type Code struct {
	ID          int64     `reform:"id,pk"`
	Code        string    `reform:"code"`
	Type        int32     `reform:"type"`
	Number      string    `reform:"number"`
	Isused      int32     `reform:"isused"`
	Userid      int64     `reform:"userid"`
	Usedatetime time.Time `reform:"usedatetime"`
}
