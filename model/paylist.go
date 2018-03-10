package model

//go:generate reform

// Paylist represents a row in paylist table.
//reform:paylist
type Paylist struct {
	ID       int64   `reform:"id,pk"`
	Userid   int64   `reform:"userid"`
	Total    string  `reform:"total"`
	Status   int32   `reform:"status"`
	Tradeno  *string `reform:"tradeno"`
	Datetime int64   `reform:"datetime"`
}
