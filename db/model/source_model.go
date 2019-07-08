package model

import "github.com/jinzhu/gorm"

type SourceModel struct {
	gorm.Model
	BookID uint
	Name   string
	Url    string
}
