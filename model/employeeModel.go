package model

import (
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	Name   string `json:"name"`
	Salary string `json:"salary"`
}

func (Employees) TableName() string {
	return "employees"
}
