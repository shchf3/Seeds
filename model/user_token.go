package model

//go:generate reform

// UserToken represents a row in user_token table.
//reform:user_token
type UserToken struct {
	ID         int32  `reform:"id,pk"`
	Token      string `reform:"token"`
	UserID     int32  `reform:"user_id"`
	CreateTime int32  `reform:"create_time"`
	ExpireTime int32  `reform:"expire_time"`
}
