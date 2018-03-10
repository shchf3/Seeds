package model

//go:generate reform

// Ticket represents a row in ticket table.
//reform:ticket
type Ticket struct {
	ID       int64  `reform:"id,pk"`
	Title    string `reform:"title"`
	Content  string `reform:"content"`
	Rootid   int64  `reform:"rootid"`
	Userid   int64  `reform:"userid"`
	Datetime int64  `reform:"datetime"`
	Status   int32  `reform:"status"`
}
