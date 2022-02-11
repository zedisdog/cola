package database

import "gorm.io/gorm"

var Configs map[string]*Conf

type Conf struct {
	Dsn         DSN
	GormSetters []func(*gorm.Config)
}

func Config(conf *Conf, name ...string) {
	if Configs == nil {
		Configs = make(map[string]*Conf, 0)
	}
	if len(name) > 0 {
		Configs[name[0]] = conf
	} else {
		Configs["default"] = conf
	}
}
