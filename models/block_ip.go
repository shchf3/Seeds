package models

type BlockIp struct {
  Id        	int64     	`json:"id";gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  NodeId    	int       	`json:"nodeid";gorm:"column:nodeid;NOT NULL;"`
  Ip        	string    	`json:"ip";gorm:"column:ip;NOT NULL;type:text;"`
  Datetime  	int64     	`json:"datetime";gorm:"column:datetime;NOT NULL;"`
}

func (BlockIp) TableName() string {
  return "blockip"
}