//+build mysql
//+build !sqlite
//+build !postgres

package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(dsn string, setOptions ...func(*options)) (*gorm.DB, error) {
	o := &options{
		config: &gorm.Config{},
	}
	for _, setOption := range setOptions {
		setOption(o)
	}
	return gorm.Open(mysql.Open(dsn), o.config)
}

type options struct {
	config *gorm.Config
}

func WithConfig(config *gorm.Config) func(o *options) {
	return func(o *options) {
		o.config = config
	}
}
