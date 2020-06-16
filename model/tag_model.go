package model

import "github.com/jinzhu/gorm"

type TagModel struct {
	gorm.Model
	Name          string // tag 名字
	Type          int    // tag 类型 一级/二级
	BookListRefer int    // 书单引用数
	BookRefer     int    //书籍引用数
}
