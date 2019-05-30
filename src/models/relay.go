package models

type Relay struct {
  Id           int64  `json:"id";gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  UserId       int64  `json:"user_id";gorm:"column:user_id;NOT NULL;"`
  SourceNodeId int64  `json:"source_node_id";gorm:"column:source_node_id;NOT NULL;"`
  DistNodeId   int64  `json:"dist_node_id";gorm:"column:dist_node_id;NOT NULL;"`
  DistIp       string `json:"dist_ip";gorm:"column:dist_ip;NOT NULL;type:text;"`
  Port         int    `json:"port";gorm:"column:port;NOT NULL;"`
  Priority     int    `json:"priority";gorm:"column:priority;NOT NULL;"`
}

func (Relay) TableName() string {
  return "relay"
}