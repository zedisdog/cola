package database

import (
	"errors"
	mocket "github.com/Selvatico/go-mocket"
	"github.com/zedisdog/cola/tools"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
)

func NewMocker() (db *gorm.DB, err error) {
	// see: https://github.com/Selvatico/go-mocket/blob/master/DOCUMENTATION.md
	// for gorm v2, see: https://github.com/Selvatico/go-mocket/issues/28
	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = true
	dialect := mysql.New(mysql.Config{
		DSN:                       "mockdb",
		DriverName:                mocket.DriverName,
		SkipInitializeWithVersion: true,
	})
	db, err = gorm.Open(dialect)
	return
}

func New(setOptions ...func(*Options)) (*gorm.DB, error) {
	o := &Options{
		config: &gorm.Config{},
	}
	for _, setOption := range setOptions {
		setOption(o)
	}
	if o.dialector != nil {
		return gorm.Open(o.dialector, o.config)
	} else {
		d, err := getDialector(o.dsn)
		if err != nil {
			panic(err)
		}
		return gorm.Open(d, o.config)
	}
}

func getDialector(dsn string) (gorm.Dialector, error) {
	reg := regexp.MustCompile(`(^\S+)://(\S+$)`)
	info := reg.FindStringSubmatch(dsn)
	if len(info) < 3 {
		return nil, errors.New("dsn is invalid")
	}

	switch info[1] {
	case "mysql":
		return mysql.Open(info[2]), nil
	case "postgres":
		return postgres.Open(info[2]), nil
	}

	return nil, errors.New("not support database type")
}

type Options struct {
	config    *gorm.Config
	dialector gorm.Dialector
	dsn       string
}

func WithConfig(config *gorm.Config) func(o *Options) {
	return func(o *Options) {
		o.config = config
	}
}

func WithDialector(d gorm.Dialector) func(o *Options) {
	return func(o *Options) {
		o.dialector = d
	}
}

func WithDsn(dsn string) func(o *Options) {
	return func(o *Options) {
		o.dsn = tools.EncodeQuery(dsn)
	}
}
