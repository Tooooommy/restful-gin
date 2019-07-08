package migrations

import (
	"github.com/go-gormigrate/gormigrate"
	"sort"
	"strconv"
)

type MigrationList []*gormigrate.Migration

func (ml MigrationList) Len() int {
	return len(ml)
}

func (ml MigrationList) Swap(i, j int) {
	ml[i], ml[j] = ml[j], ml[i]
}
func (ml MigrationList) Less(i, j int) bool {
	idi, _ := strconv.Atoi(ml[i].ID)
	idj, _ := strconv.Atoi(ml[j].ID)
	return idi < idj
}

var MigrationTpl = `package migrations

import (
	"github.com/go-gormigrate/gormigrate"
	"github.com/jinzhu/gorm"
)


// previous id %s
func Migration%s() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "%s",
		Migrate: func(db *gorm.DB) error {
			// migrate code
			return nil
		},
		Rollback: func(db *gorm.DB) error {
			// rollback code
			return nil
		},
	}
}`

func GetMigrations() MigrationList {
	var migrations = MigrationList{
		Migrationtest(),
		Migrationtest1(),
	}
	sort.Sort(migrations)
	return migrations
}
