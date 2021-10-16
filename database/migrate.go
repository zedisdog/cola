package database

import (
	"embed"
	"github.com/golang-migrate/migrate/v4"
	migrate2 "github.com/zedisdog/cola/cmd/cola/migrate"
	"github.com/zedisdog/cola/tools"
)

// AutoMigrate 自动迁移
func AutoMigrate(migrations embed.FS, dsn string) {
	m, err := migrate.NewWithSourceInstance(
		"",
		migrate2.NewEmbed(migrations),
		tools.EncodeQuery(dsn),
	)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		panic(err)
	}
}
