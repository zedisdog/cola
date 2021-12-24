package migrate

import (
	_ "github.com/golang-migrate/migrate/v4/database/mysql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/zedisdog/cola/tools"
)

// AutoMigrate 自动迁移
func AutoMigrate(dsn string) {
	m, err := migrate.NewWithSourceInstance(
		"",
		EmbedDriver,
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
