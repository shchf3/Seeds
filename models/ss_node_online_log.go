package models

type SsNodeOnlineLog struct {
  Id         int `gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  NodeId     int `gorm:"column:node_id;NOT NULL;"`
  OnlineUser int `gorm:"column:online_user;NOT NULL;"`
  LogTime    int `gorm:"column:log_time;NOT NULL;"`
}

func (SsNodeOnlineLog) TableName() string {
  return "ss_node_online_log"
}