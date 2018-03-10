package model

//go:generate reform

// UserTrafficLog represents a row in user_traffic_log table.
//reform:user_traffic_log
type UserTrafficLog struct {
	ID      int32   `reform:"id,pk"`
	UserID  int32   `reform:"user_id"`
	U       int64   `reform:"u"`
	D       int64   `reform:"d"`
	NodeID  int32   `reform:"node_id"`
	Rate    float32 `reform:"rate"`
	Traffic string  `reform:"traffic"`
	LogTime int32   `reform:"log_time"`
}
