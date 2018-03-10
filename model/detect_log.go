package model

//go:generate reform

// DetectLog represents a row in detect_log table.
//reform:detect_log
type DetectLog struct {
	ID       int64 `reform:"id,pk"`
	UserID   int64 `reform:"user_id"`
	ListID   int64 `reform:"list_id"`
	Datetime int64 `reform:"datetime"`
	NodeID   int32 `reform:"node_id"`
}
