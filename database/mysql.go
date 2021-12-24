//go:build !gorm

package database

import (
	"database/sql"
	"github.com/zedisdog/cola/tools"
	"strings"
)

var DB *sql.DB

func InitDB(dsn string) {
	if DB == nil {
		var err error
		DB, err = sql.Open(
			"mysql",
			strings.Replace(tools.EncodeQuery(dsn), "mysql://", "", 1),
		)
		if err != nil {
			panic(err)
		}
	}
}
