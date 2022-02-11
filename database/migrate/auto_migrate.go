package migrate

import (
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/zedisdog/cola/database"

	"github.com/golang-migrate/migrate/v4"
)

// AutoMigrate 自动迁移
func AutoMigrate(name ...string) {
	var n string
	if len(name) > 0 {
		n = name[0]
	} else {
		n = "default"
	}
	m, err := migrate.NewWithSourceInstance(
		"",
		EmbedDriver,
		database.Configs[n].Dsn.Encode(),
	)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && err.Error() != "no change" {
		panic(err)
	}
}
