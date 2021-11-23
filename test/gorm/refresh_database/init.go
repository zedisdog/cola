package refresh_database

import (
	"errors"
	"github.com/zedisdog/cola/database"
)

func init() {
	db := database.Instance()
	if db == nil {
		panic(errors.New("no instance of gorm.DB"))
	}
	database.SetInstance("default", db.Begin())
}
