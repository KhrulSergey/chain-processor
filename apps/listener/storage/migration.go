package storage

import (
	"embed"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var (
	//go:embed migrations/*.sql
	migrations embed.FS
)

func MigrateGorm(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID:      "V202309041652540  create tables",
			Migrate: migrateFromFile("migrations/V202309041652540__create_tables.sql"),
		},
	})

	return m.Migrate()
}

func migrateFromFile(migrationFile string) gormigrate.MigrateFunc {
	return func(tx *gorm.DB) error {
		sql, err := migrations.ReadFile(migrationFile)
		if err != nil {
			return err
		}
		return tx.Exec(string(sql)).Error
	}
}
