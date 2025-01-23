package storage

import (
	"embed"
	"strings"

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
			ID:      "V202501191422540  create trade",
			Migrate: migrateFromFile("migrations/V202501191422540__create_trade.sql"),
		},
		{
			ID:      "V202501201652540  create user",
			Migrate: migrateFromFile("migrations/V202501201652540__create_user.sql"),
		},
		{
			ID:      "V202501201952540  create tables",
			Migrate: migrateFromFile("migrations/V202501201952540__create_auth.sql"),
		},
	})

	return m.Migrate()
}

func migrateFromFile(migrationFile string) gormigrate.MigrateFunc {
	return func(tx *gorm.DB) error {
		sqlContent, err := migrations.ReadFile(migrationFile)
		if err != nil {
			return err
		}
		queries := strings.Split(string(sqlContent), ";")
		for _, query := range queries {
			trimmedQuery := strings.TrimSpace(query)
			if trimmedQuery != "" {
				if err := tx.Exec(trimmedQuery).Error; err != nil {
					return err
				}
			}
		}
		return nil
	}
}
