package model

import (
	"time"
)

//go:generate reform

// User represents a row in user table.
//reform:user
type User struct {
	ID                 int32     `reform:"id,pk"`
	UserName           string    `reform:"user_name"`
	Email              string    `reform:"email"`
	Pass               string    `reform:"pass"`
	Passwd             string    `reform:"passwd"`
	T                  int32     `reform:"t"`
	U                  int64     `reform:"u"`
	D                  int64     `reform:"d"`
	Plan               string    `reform:"plan"`
	TransferEnable     int64     `reform:"transfer_enable"`
	Port               int32     `reform:"port"`
	Switch             int8      `reform:"switch"`
	Enable             int8      `reform:"enable"`
	Type               int8      `reform:"type"`
	LastGetGiftTime    int32     `reform:"last_get_gift_time"`
	LastCheckInTime    int32     `reform:"last_check_in_time"`
	LastRestPassTime   int32     `reform:"last_rest_pass_time"`
	RegDate            time.Time `reform:"reg_date"`
	InviteNum          int32     `reform:"invite_num"`
	Money              string    `reform:"money"`
	RefBy              int32     `reform:"ref_by"`
	ExpireTime         int32     `reform:"expire_time"`
	Method             string    `reform:"method"`
	IsEmailVerify      int8      `reform:"is_email_verify"`
	RegIp              string    `reform:"reg_ip"`
	NodeSpeedlimit     string    `reform:"node_speedlimit"`
	NodeConnector      int32     `reform:"node_connector"`
	IsAdmin            int32     `reform:"is_admin"`
	ImType             *int32    `reform:"im_type"`
	ImValue            *string   `reform:"im_value"`
	LastDayT           int64     `reform:"last_day_t"`
	SendDailyMail      int32     `reform:"sendDailyMail"`
	Class              int32     `reform:"class"`
	ClassExpire        time.Time `reform:"class_expire"`
	ExpireIn           time.Time `reform:"expire_in"`
	Theme              string    `reform:"theme"`
	GaToken            string    `reform:"ga_token"`
	GaEnable           int32     `reform:"ga_enable"`
	Pac                *string   `reform:"pac"`
	Remark             *string   `reform:"remark"`
	NodeGroup          int32     `reform:"node_group"`
	AutoResetDay       int32     `reform:"auto_reset_day"`
	AutoResetBandwidth string    `reform:"auto_reset_bandwidth"`
	Protocol           *string   `reform:"protocol"`
	ProtocolParam      *string   `reform:"protocol_param"`
	Obfs               *string   `reform:"obfs"`
	ObfsParam          *string   `reform:"obfs_param"`
	ForbiddenIp        *string   `reform:"forbidden_ip"`
	ForbiddenPort      *string   `reform:"forbidden_port"`
	DisconnectIp       *string   `reform:"disconnect_ip"`
	IsHide             int32     `reform:"is_hide"`
	IsMultiUser        int32     `reform:"is_multi_user"`
	TelegramID         *int64    `reform:"telegram_id"`
}
