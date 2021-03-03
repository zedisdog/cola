package migrate

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"github.com/zedisdog/cola/pather"
	"os"
	"os/exec"
	"sync"
)

var instance *migrate.Migrate
var once sync.Once

func GetInstance() (*migrate.Migrate, error) {
	var err error
	once.Do(func() {
		path := fmt.Sprintf("file://%s", viper.GetString("migratePath"))
		instance, err = migrate.New(path, viper.GetString("dsn"))
	})
	return instance, err
}

func Create(fileName string, ext string, prefixFormat string) error {
	p := pather.New(viper.GetString("root"))
	cmd := exec.Command(
		"migrate",
		"create",
		"-format",
		prefixFormat,
		"-ext",
		ext,
		fileName,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = p.Gen("internal/database/migrations")
	return cmd.Run()
}
