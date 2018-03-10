package model

import (
	"time"
)

//go:generate reform

// SsInviteCode represents a row in ss_invite_code table.
//reform:ss_invite_code
type SsInviteCode struct {
	ID        int32     `reform:"id,pk"`
	Code      string    `reform:"code"`
	UserID    int32     `reform:"user_id"`
	CreatedAt time.Time `reform:"created_at"`
	UpdatedAt time.Time `reform:"updated_at"`
}
