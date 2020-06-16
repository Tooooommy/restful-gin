package model

import (
	"fmt"
	"restful-gin/db"
	"testing"
)

var DefaultPath = "/home/tommy/go/src/CrownDaisy_GOGIN/app.ini"

func TestAccountModel(t *testing.T) {
	db := db.GetGormAuto()
	//am := &AccountModel{
	//	Avatar:    "https://huangzijian.com/avatar.jpg",
	//	Name:      "tommy-huang",
	//	Email:     "tommy-huang@164.com",
	//	Password:  "12736hkaf",
	//	SessionID: "tony",
	//	Status:    false,
	//}

	//// == create
	//if db.NewRecord(am) {
	//	fmt.Println("account model not exist.")
	//}
	//db.Create(am)
	//if db.NewRecord(am) {
	//	fmt.Println("account model not exist")
	//}
	//
	//bl := &BookListModel{
	//	Name:        "BoooooooKList",
	//	Description: "你好，我是你爸爸",
	//	OwnerID:     2,
	//	Tags:        "tages",
	//	TagIds:      "tags.",
	//}
	//db.Create(bl)
	//
	//user := &AccountModel{}
	//user.ID = 2
	//blm := &BookListModel{}
	//db.Model(user).Related(blm, "owner_id")
	//fmt.Printf("%v\n", blm)

	user := &AccountModel{}
	user.Name = "tommy-huang"
	if db.NewRecord(user) {
		fmt.Printf("new record")
	}
	//blm := &BookListModel{}
	//db.Model(user).Related(blm, "BookLists")
	//db.Preload("BookLists").First(user)
	//db.First(user)
	//fmt.Printf("%v\n", blm)
	//fmt.Printf("%v\n", user.BookLists)
	//db.Where(user).First(user)
	//fmt.Printf("%+v\n", user)
}
