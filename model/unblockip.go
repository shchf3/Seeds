package model

//go:generate reform

// Unblockip represents a row in unblockip table.
//reform:unblockip
type Unblockip struct {
	ID       int64  `reform:"id,pk"`
	Ip       string `reform:"ip"`
	Datetime int64  `reform:"datetime"`
	Userid   int64  `reform:"userid"`
}
