package model

//go:generate reform

// Blockip represents a row in blockip table.
//reform:blockip
type Blockip struct {
	ID       int64  `reform:"id,pk"`
	Nodeid   int32  `reform:"nodeid"`
	Ip       string `reform:"ip"`
	Datetime int64  `reform:"datetime"`
}
