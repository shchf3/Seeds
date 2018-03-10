package model

//go:generate reform

// SsNodeOnlineLog represents a row in ss_node_online_log table.
//reform:ss_node_online_log
type SsNodeOnlineLog struct {
	ID         int32 `reform:"id,pk"`
	NodeID     int32 `reform:"node_id"`
	OnlineUser int32 `reform:"online_user"`
	LogTime    int32 `reform:"log_time"`
}
