package entity

import "github.com/jinzhu/gorm"

type DateClientInfo struct {
	gorm.Model
	Host string
	Name string
	Date string
}
