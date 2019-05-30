package models

type SsNodeInfo struct {
  Id      int     `gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  NodeId  int     `gorm:"column:node_id;NOT NULL;"`
  Uptime  float64 `gorm:"column:uptime;NOT NULL;"`
  Load    string  `gorm:"column:load;NOT NULL;"`
  LogTime int     `gorm:"column:log_time;NOT NULL;"`
}

func (SsNodeInfo) TableName() string {
  return "ss_node_info"
}