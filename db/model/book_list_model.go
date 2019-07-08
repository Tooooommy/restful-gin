package model

import "github.com/jinzhu/gorm"

type BookListModel struct {
	gorm.Model
	Name        string         // 书单名字
	Description string         // 书单简介
	OwnerID     int            // 书单创建人
	Tags        string         // 标签列表
	TagIds      string         // 标签列表ID
	Books       []BookModel    `gorm:"many2many:book_in_list"`          // 书单列表 many to many
	Followers   []AccountModel `gorm:"many2many:user_follow_book_list"` //书单关注者 many to many
	Comments    []CommentModel `gorm:"ForeignKey:AccountID"`            // 书单评论 1 to many
}
