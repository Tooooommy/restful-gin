package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"restful-gin/config"
	"strings"
	"time"
)

var gdb *gorm.DB

func InitMysqlDB() (err error) {
	// 数据库表默认名字
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		var tableName = defaultTableName
		if strings.HasSuffix(defaultTableName, "_models") {
			tableNameParts := strings.Split(defaultTableName, "_")
			tableName = strings.Join(tableNameParts[:len(tableNameParts)-1], "_") + "s"
		}
		return tableName
	}

	cfg := config.Get().Mysql
	if cfg.Charset == "" {
		cfg.Charset = "utf8mb4"
	}
	if cfg.Loc == "" {
		cfg.Loc = "UTC"
	}

	if cfg.ConnectionTimeout == 0 {
		cfg.ConnectionTimeout = 10
	}
	dbSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&loc=%s&parseTime=true&timeout=%ds",
		cfg.Username, cfg.Password, cfg.Host, cfg.Schema, cfg.Charset, cfg.Loc, cfg.ConnectionTimeout)
	gdb, err = gorm.Open("mysql", dbSource)
	if err != nil {
		return err
	}
	gdb.DB().SetConnMaxLifetime(time.Duration(cfg.MaxConnLifetime) * time.Second)
	gdb.DB().SetMaxIdleConns(cfg.MaxIdleConn)
	gdb.DB().SetMaxOpenConns(cfg.MaxOpenConn)
	go func() {
		for {
			time.Sleep(10 * time.Second)
			err := gdb.DB().Ping()
			if err != nil {
				fmt.Printf("database ping error: %+v\n", err)
			}
		}
	}()
	return
}

func GetGormAuto() *gorm.DB {
	if gdb == nil {
		if err := InitMysqlDB(); err != nil {
			panic(nil)
		}
	}
	return gdb
}
