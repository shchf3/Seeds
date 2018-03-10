package models

type UnblockIp struct {
  Id        	int64     	`json:"id";gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  Ip        	string    	`json:"ip";gorm:"column:ip;NOT NULL;type:text;"`
  Datetime  	int64     	`json:"datetime";gorm:"column:datetime;NOT NULL;"`
  Userid    	int64     	`json:"userid";gorm:"column:userid;NOT NULL;"`
}

func (UnblockIp) TableName() string {
  return "unblockip"
}