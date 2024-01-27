package entity

import "github.com/jinzhu/gorm"

type ClientInfo struct {
	gorm.Model
	Host     string
	Name     string
	DateTime string
}
