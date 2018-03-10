package model

//go:generate reform

// EmailVerify represents a row in email_verify table.
//reform:email_verify
type EmailVerify struct {
	ID       int64  `reform:"id,pk"`
	Email    string `reform:"email"`
	Ip       string `reform:"ip"`
	Code     string `reform:"code"`
	ExpireIn int64  `reform:"expire_in"`
}
