package models

type Auto struct {
  Id        	int64     	`json:"id";gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  Type      	int       	`json:"type";gorm:"column:type;NOT NULL;"`
  Value     	string    	`json:"value";gorm:"column:value;NOT NULL;type:longtext;"`
  Sign      	string    	`json:"sign";gorm:"column:sign;NOT NULL;type:longtext;"`
  Datetime  	int64     	`json:"datetime";gorm:"column:datetime;NOT NULL;"`
}

func (Auto) TableName() string {
  return "auto"
}