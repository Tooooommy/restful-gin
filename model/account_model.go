package model

import (
	"github.com/jinzhu/gorm"
	"restful-gin/db"
)

type AccountModel struct {
	*gorm.Model
	UniqueId  string          //唯一id
	Avatar    string          // 用户头像
	Name      string          // 用户名
	Email     string          // 邮箱
	Password  string          // 密码
	SessionId string          // token session id
	Status    bool            // 用户状态
	BookLists []BookListModel `gorm:"ForeignKey:OwnerID"`   // 用户拥有的书单 1 to many
	Rates     []RateModel     `gorm:"ForeignKey:AccountID"` // 评价的书籍 1 to many
}

func (m *AccountModel) GetTable() *gorm.DB {
	return db.GetGormAuto().Table("accounts")
}

func (m *AccountModel) GetDB() *gorm.DB {
	return db.GetGormAuto()
}

func (m *AccountModel) Exist() {
	m.GetDB().NewRecord(m)
}

func (m *AccountModel) GetAccount() (model *AccountModel, exist bool) {
	exist = !m.GetDB().Where(m).First(model).NewRecord(m)
	return
}

func (m *AccountModel) GetAccountByName(name string) (model *AccountModel, exist bool) {
	exist = !m.GetDB().Where(&AccountModel{Name: name}).First(model).NewRecord(m)
	return
}

func (m *AccountModel) GetAccountByEmail(email string) (model *AccountModel, exist bool) {
	exist = !m.GetDB().Where(&AccountModel{Email: email}).First(model).NewRecord(m)
	return
}
