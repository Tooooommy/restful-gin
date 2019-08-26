package main

import (
	"flag"
	"fmt"
	"github.com/go-gormigrate/gormigrate"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"restful-gin/config"
	"restful-gin/db/database/migrations"
	"strings"
	"time"
)

// 生产一个migration文件
// gen
// migrate / migrateTo
// rollback / rollbackTo
var ms = flag.Bool("ms", false, "list migrations")
var g = flag.String("g", "", "generate new migrations")
var m = flag.Bool("m", false, "migrate all migration")
var mt = flag.String("mt", "", "migrate to migration")
var r = flag.Bool("r", false, "rollback last migration")
var rt = flag.String("rt", "", "rollback to migration")

func InitMigrate(msg migrations.MigrationList) (*gormigrate.Gormigrate, error) {
	cfg := config.Get().Mysql
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&loc=%s&parseTime=true&timeout=%ds",
		cfg.Username, cfg.Password, cfg.Host, cfg.Schema, cfg.Charset, cfg.Loc, cfg.ConnectionTimeout)
	db, err := gorm.Open("mysql", source)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return gormigrate.New(db, gormigrate.DefaultOptions, msg), nil
}

func main() {
	mgs := migrations.GetMigrations()
	migrate, err := InitMigrate(mgs)
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	if *ms {
		CmdListMigrations(mgs)
		return
	}
	if *g != "" {
		CmdGenerateMigration(mgs)
		return
	}
	if *m {
		CmdMigrate(migrate)
		return
	}
	if *mt != "" {
		CmdMigrateTo(migrate)
	}
	if *r {
		CmdRollback(migrate)
	}
	if *rt != "" {
		CmdRollbackTo(migrate)
	}
}

func CmdListMigrations(mgs migrations.MigrationList) {
	var ids = make([]string, 0)
	for i := len(mgs) - 1; i >= 0; i-- {
		ids = append(ids, mgs[i].ID)
	}
	if len(ids) > 0 {
		fmt.Printf("%s\n", strings.Join(ids, "\n"))
	} else {
		fmt.Println("Not Found Migrations")
	}
	fmt.Printf("success list migrations\n")
}

func CmdGenerateMigration(mgs migrations.MigrationList) {
	fileInfo, err := os.Stat("migrations")
	if err != nil {
		err = os.Mkdir("./migrations", 666)
		panic(err)
	}
	if !fileInfo.IsDir() {
		panic("migrations is not a dir")
	}
	_ = os.Chdir("migrations")

	var previousId = "0"
	if len(mgs) > 0 {
		previousId = mgs[len(mgs)-1].ID
	}
	now := time.Now().Format("200612150405")
	fileName := fmt.Sprintf("%s_%s_%s.go", "migration", *g, now)

	tpl := fmt.Sprintf(migrations.MigrationTpl, previousId, *g, now)
	file, err := os.Create(fileName)
	if err != nil {
		panic("create migration failed")
	}
	defer file.Close()
	_, err = file.WriteString(tpl)
	if err != nil {
		panic("write migration tmpl failed")
	}
	fmt.Printf("success generate migration file\n")
}

func CmdMigrate(m *gormigrate.Gormigrate) {
	err := m.Migrate()
	if err != nil {
		panic(err)
	}

	fmt.Printf("success migrate all\n")
}

func CmdMigrateTo(m *gormigrate.Gormigrate) {
	err := m.MigrateTo(*mt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("success migrate to %s\n", *mt)
}

func CmdRollback(m *gormigrate.Gormigrate) {
	err := m.RollbackLast()
	if err != nil {
		panic(err)
	}

	fmt.Printf("success rollback last\n")
}

func CmdRollbackTo(m *gormigrate.Gormigrate) {
	err := m.RollbackTo(*rt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("success rollback to %s\n", *rt)
}
