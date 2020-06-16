package model

import "github.com/jinzhu/gorm"

type RateModel struct {
	gorm.Model
	BookID    int
	AccountID int
	Rate      int
}
