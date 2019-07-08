package model

import (
	"github.com/jinzhu/gorm"
)

type CommentModel struct {
	gorm.Model
	BookListID int    // 评论的书单
	AccountID  int    // 评论用户ID
	Content    string // 评论内容
	ReplyID    int    // 回复的CommentID, 没有就是父评论
}
