package model

import (
	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
)

const DNS = "sqlserver://developer:d123@90.0.0.110?database=BasicTraining"

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
