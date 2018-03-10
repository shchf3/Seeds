package models

type AliveIp struct {
  Id        	int64     	`gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  Nodeid    	int       	`gorm:"column:nodeid;NOT NULL;"`
  Userid    	int       	`gorm:"column:userid;NOT NULL;"`
  Ip        	string    	`gorm:"column:ip;NOT NULL;type:text;"`
  Datetime  	int64     	`gorm:"column:datetime;NOT NULL;"`
}

func (AliveIp) TableName() string {
  return "alive_ip"
}