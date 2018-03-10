package model

//go:generate reform

// RadiusBan represents a row in radius_ban table.
//reform:radius_ban
type RadiusBan struct {
	ID     int32 `reform:"id,pk"`
	Userid int32 `reform:"userid"`
}
