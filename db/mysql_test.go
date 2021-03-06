package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"restful-gin/config"
	"testing"
)

func TestConnect(t *testing.T) {
	cfg := config.Get().Mysql
	fmt.Println(cfg)
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&loc=%s&parseTime=true&timeout=%ds",
		cfg.Username, cfg.Password, cfg.Host, cfg.Schema, cfg.Charset, cfg.Loc, cfg.ConnectionTimeout)
	db, err := gorm.Open("mysql", source)
	if err != nil {
		t.Errorf("open mysql connnect error: %v", err)
		return
	}
	defer db.Close()
	if err := db.DB().Ping(); err != nil {
		t.Errorf("ping error: %v", err)
	}
}

type TestAccount struct {
	gorm.Model
	TestName     string
	TestPassword string
}

func TestGDB(t *testing.T) {
	err := InitMysqlDB()
	if err != nil {
		t.Errorf("test init mysql error: %v", err)
	}

	db := GetGormAuto()
	defer db.Close()

	// == ping
	if db.DB().Ping() != nil {
		t.Error("db ping error")
	}

	// == create table
	db.CreateTable(&TestAccount{})
	if !db.HasTable("test_accounts") {
		t.Errorf("create table test_accounts error")
	}
	db.DropTable("test_accounts")
	if db.HasTable("test_accounts") {
		t.Errorf("drop table test_accounts error")
	}
}
