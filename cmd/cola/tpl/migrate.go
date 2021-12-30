package tpl

const MigrateTemp = `package migration

import (
	"embed"
	"github.com/zedisdog/cola/database/migrate"
)

//go:embed *.sql
var fs embed.FS

func Register(driver *migrate.EDriver) {
	driver.Add(&fs)
}
`
