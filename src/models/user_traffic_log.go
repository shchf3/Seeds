package models

type UserTrafficLog struct {
  Id      int     `gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  UserId  int     `gorm:"column:user_id;NOT NULL;"`
  U       int64   `gorm:"column:u;NOT NULL;"`
  D       int64   `gorm:"column:d;NOT NULL;"`
  NodeId  int     `gorm:"column:node_id;NOT NULL;"`
  Rate    float64 `gorm:"column:rate;NOT NULL;"`
  Traffic string  `gorm:"column:traffic;NOT NULL;"`
  LogTime int     `gorm:"column:log_time;NOT NULL;"`
}

func (UserTrafficLog) TableName() string {
  return "user_traffic_log"
}