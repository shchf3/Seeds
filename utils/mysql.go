package utils

import (
	"github.com/jinzhu/gorm"
	"sync"
	"log"
	"Seeds/models"
)

type singletonMySQL struct {
	Database *gorm.DB
}

var instance *singletonMySQL
var lock sync.Once

func GetMySQLInstance() *singletonMySQL {
	lock.Do(func() {
		// Init MySQL Instance
		db, err := gorm.Open("mysql", GetDataUrl())
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		db.LogMode(true)
		db.AutoMigrate(&models.SsNode{})
		db.AutoMigrate(&models.User{})

		instance = &singletonMySQL{
			Database: db,
		}
	})
	return instance
}