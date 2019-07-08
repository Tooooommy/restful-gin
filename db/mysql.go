package db

import (
	"CrownDaisy_GOGIN/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"time"
)

var GDB *gorm.DB

func InitMysqlDB() (err error) {
	// 名字
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
	GDB, err = gorm.Open("mysql", dbSource)
	if err != nil {
		return err
	}
	GDB.DB().SetConnMaxLifetime(time.Duration(cfg.MaxConnLifetime) * time.Second)
	GDB.DB().SetMaxIdleConns(cfg.MaxIdleConns)
	GDB.DB().SetMaxOpenConns(cfg.MaxOpenConns)
	go func() {
		for {
			time.Sleep(10 * time.Second)
			err := GDB.DB().Ping()
			if err != nil {
				fmt.Printf("database ping error: %+v\n", err)
			}
		}
	}()
	return
}

func GetMysqlDB() *gorm.DB {
	if GDB == nil {
		if err := InitMysqlDB(); err != nil {
			panic(nil)
		}
	}
	return GDB
}
