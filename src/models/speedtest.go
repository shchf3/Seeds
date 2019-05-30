package models

type SpeedTest struct {
  Id        	    int64     	`json:"id";gorm:"column:id;NOT NULL;PRIMARY KEY;"`
  NodeId    	    int       	`json:"nodeid";gorm:"column:nodeid;NOT NULL;"`
  DateTime  	    int64     	`json:"datetime";gorm:"column:datetime;NOT NULL;"`
  TelecomePing	    string    	`json:"telecomping";gorm:"column:telecomping;NOT NULL;type:text;"`
  TelecomeUpload	string    	`json:"telecomeupload";gorm:"column:telecomeupload;NOT NULL;type:text;"`
  TelecomeDownload	string    	`json:"telecomedownload";gorm:"column:telecomedownload;NOT NULL;type:text;"`
  UnicomPing	    string    	`json:"unicomping";gorm:"column:unicomping;NOT NULL;type:text;"`
  UnicomUpload	    string    	`json:"unicomupload";gorm:"column:unicomupload;NOT NULL;type:text;"`
  UnicomDownload	string    	`json:"unicomdownload";gorm:"column:unicomdownload;NOT NULL;type:text;"`
  CmccPing  	    string    	`json:"cmccping";gorm:"column:cmccping;NOT NULL;type:text;"`
  CmccUpload	    string    	`json:"cmccupload";gorm:"column:cmccupload;NOT NULL;type:text;"`
  CmccDownload	    string    	`json:"cmccdownload";gorm:"column:cmccdownload;NOT NULL;type:text;"`
}

func (SpeedTest) TableName() string {
  return "speedtest"
}