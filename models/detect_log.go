package models

type DetectLog struct {
  Id       int64 `gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  UserId   int64 `gorm:"column:user_id;NOT NULL;"`
  ListId   int64 `gorm:"column:list_id;NOT NULL;"`
  Datetime int64 `gorm:"column:datetime;NOT NULL;"`
  NodeId   int   `gorm:"column:node_id;NOT NULL;"`
}

func (DetectLog) TableName() string {
  return "detect_log"
}