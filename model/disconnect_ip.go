package model

//go:generate reform

// DisconnectIp represents a row in disconnect_ip table.
//reform:disconnect_ip
type DisconnectIp struct {
	ID       int64  `reform:"id,pk"`
	Userid   int64  `reform:"userid"`
	Ip       string `reform:"ip"`
	Datetime int64  `reform:"datetime"`
}
