package model

//go:generate reform

// Relay represents a row in relay table.
//reform:relay
type Relay struct {
	ID           int64  `reform:"id,pk"`
	UserID       int64  `reform:"user_id"`
	SourceNodeID int64  `reform:"source_node_id"`
	DistNodeID   int64  `reform:"dist_node_id"`
	DistIp       string `reform:"dist_ip"`
	Port         int32  `reform:"port"`
	Priority     int32  `reform:"priority"`
}
