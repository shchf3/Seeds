package model

//go:generate reform

// Bought represents a row in bought table.
//reform:bought
type Bought struct {
	ID       int64  `reform:"id,pk"`
	Userid   int64  `reform:"userid"`
	Shopid   int64  `reform:"shopid"`
	Datetime int64  `reform:"datetime"`
	Renew    int64  `reform:"renew"`
	Coupon   string `reform:"coupon"`
	Price    string `reform:"price"`
}
