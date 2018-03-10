package model

//go:generate reform

// LoginIp represents a row in login_ip table.
//reform:login_ip
type LoginIp struct {
	ID       int64  `reform:"id,pk"`
	Userid   int64  `reform:"userid"`
	Ip       string `reform:"ip"`
	Datetime int64  `reform:"datetime"`
	Type     int32  `reform:"type"`
}
