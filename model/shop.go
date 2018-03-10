package model

//go:generate reform

// Shop represents a row in shop table.
//reform:shop
type Shop struct {
	ID                 int64  `reform:"id,pk"`
	Name               string `reform:"name"`
	Price              string `reform:"price"`
	Content            string `reform:"content"`
	AutoRenew          int32  `reform:"auto_renew"`
	AutoResetBandwidth int32  `reform:"auto_reset_bandwidth"`
	Status             int32  `reform:"status"`
}
