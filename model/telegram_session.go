package model

//go:generate reform

// TelegramSession represents a row in telegram_session table.
//reform:telegram_session
type TelegramSession struct {
	ID             int64  `reform:"id,pk"`
	UserID         int64  `reform:"user_id"`
	Type           int32  `reform:"type"`
	SessionContent string `reform:"session_content"`
	Datetime       int64  `reform:"datetime"`
}
