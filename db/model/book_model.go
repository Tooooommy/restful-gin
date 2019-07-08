package model

import "github.com/jinzhu/gorm"

type BookModel struct {
	gorm.Model
	Cover         string        // 封面
	Author        string        // 作者
	Name          string        // 书名
	Description   string        // 简介
	FromWebSite   string        // 来自
	NewChapter    string        // 最新章节
	WordCount     string        //字数
	ChapterCount  int           // 章节数目
	BookListRefer int           // 引用数
	Tags          string        // 标签列表
	TagIds        string        // 标签ID列表
	Rates         []RateModel   `gorm:"ForeignKey:BookID"` // 评价 1 to many
	Sources       []SourceModel `gorm:"ForeignKey:BookID"` // 来源 1 对 many
}
