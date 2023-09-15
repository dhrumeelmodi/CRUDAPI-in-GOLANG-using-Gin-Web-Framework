package model

import (
	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
)
//ConnectionString of your Database

const DNS = "" 

func Database() (*gorm.DB, error) {
	db, err := gorm.Open(sqlserver.Open(DNS), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Employees{}); err != nil {
		panic(err)
	}

	return db, err

}
