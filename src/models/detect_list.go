package models

type DetectList struct {
  Id        	int64     	`json:"id";gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  Name      	string    	`json:"name";gorm:"column:name;NOT NULL;type:longtext;"`
  Text      	string    	`json:"text";gorm:"column:text;NOT NULL;type:longtext;"`
  Regex     	string    	`json:"regex";gorm:"column:regex;NOT NULL;type:longtext;"`
  Type      	int       	`json:"type";gorm:"column:type;NOT NULL;"`
}

func (DetectList) TableName() string {
  return "detect_list"
}