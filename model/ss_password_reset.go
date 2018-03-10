package model

//go:generate reform

// SsPasswordReset represents a row in ss_password_reset table.
//reform:ss_password_reset
type SsPasswordReset struct {
	ID         int32  `reform:"id,pk"`
	Email      string `reform:"email"`
	Token      string `reform:"token"`
	InitTime   int32  `reform:"init_time"`
	ExpireTime int32  `reform:"expire_time"`
}
