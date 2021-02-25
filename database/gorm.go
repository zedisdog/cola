package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/url"
	"regexp"
)

func New(setOptions ...func(*options)) (*gorm.DB, error) {
	o := &options{
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

type options struct {
	config    *gorm.Config
	dialector gorm.Dialector
	dsn       string
}

func WithConfig(config *gorm.Config) func(o *options) {
	return func(o *options) {
		o.config = config
	}
}

func WithDialector(d gorm.Dialector) func(o *options) {
	return func(o *options) {
		o.dialector = d
	}
}

func WithDsn(dsn string) func(o *options) {
	return func(o *options) {
		u, err := url.Parse(dsn)
		if err != nil {
			panic(err)
		}
		d := fmt.Sprintf("%s://%s@%s%s?%s", u.Scheme, u.User, u.Host, u.Path, u.Query().Encode())
		o.dsn = d
	}
}
