package models

type AliveIp struct {
  Id        	int64     	`gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  NodeId    	int       	`gorm:"column:nodeid;NOT NULL;"`
  UserId    	int       	`gorm:"column:userid;NOT NULL;"`
  Ip        	string    	`gorm:"column:ip;NOT NULL;type:text;"`
  Datetime  	int64     	`gorm:"column:datetime;NOT NULL;"`
}

func (AliveIp) TableName() string {
  return "alive_ip"
}