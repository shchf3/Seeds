package model

//go:generate reform

// Auto represents a row in auto table.
//reform:auto
type Auto struct {
	ID       int64  `reform:"id,pk"`
	Type     int32  `reform:"type"`
	Value    string `reform:"value"`
	Sign     string `reform:"sign"`
	Datetime int64  `reform:"datetime"`
}
