package model

//go:generate reform

// Payback represents a row in payback table.
//reform:payback
type Payback struct {
	ID       int64  `reform:"id,pk"`
	Total    string `reform:"total"`
	Userid   int64  `reform:"userid"`
	RefBy    int64  `reform:"ref_by"`
	RefGet   string `reform:"ref_get"`
	Datetime int64  `reform:"datetime"`
}
