package migrate

import (
	"embed"
	"os"
	"os/exec"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"github.com/zedisdog/cola/pather"
	"github.com/zedisdog/cola/tools"
)

var instance *migrate.Migrate
var once sync.Once

func GetInstance() (*migrate.Migrate, error) {
	var err error
	once.Do(func() {
		f := viper.Get("migrations").(embed.FS)
		instance, err = migrate.NewWithSourceInstance(
			"",
			NewEnbed(f),
			tools.EncodeQuery(viper.GetString("database.dsn")),
		)
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
