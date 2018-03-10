package model

//go:generate reform

// Coupon represents a row in coupon table.
//reform:coupon
type Coupon struct {
	ID      int64  `reform:"id,pk"`
	Code    string `reform:"code"`
	Onetime int32  `reform:"onetime"`
	Expire  int64  `reform:"expire"`
	Shop    string `reform:"shop"`
	Credit  int32  `reform:"credit"`
}
