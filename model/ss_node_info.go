package model

//go:generate reform

// SsNodeInfo represents a row in ss_node_info table.
//reform:ss_node_info
type SsNodeInfo struct {
	ID      int32   `reform:"id,pk"`
	NodeID  int32   `reform:"node_id"`
	Uptime  float32 `reform:"uptime"`
	Load    string  `reform:"load"`
	LogTime int32   `reform:"log_time"`
}
