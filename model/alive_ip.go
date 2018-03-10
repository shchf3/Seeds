package model

//go:generate reform

// AliveIp represents a row in alive_ip table.
//reform:alive_ip
type AliveIp struct {
	ID       int64  `reform:"id,pk"`
	Nodeid   int32  `reform:"nodeid"`
	Userid   int32  `reform:"userid"`
	Ip       string `reform:"ip"`
	Datetime int64  `reform:"datetime"`
}
