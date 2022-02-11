package gorm

import (
	"fmt"
	"github.com/zedisdog/cola/database"
	"github.com/zedisdog/cola/errx"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbs map[string]*gorm.DB

func Instance(name ...string) *gorm.DB {
	var n string
	if len(name) > 0 {
		n = name[0]
	} else {
		n = "default"
	}
	if db, ok := dbs[n]; !ok {
		if c, ok := database.Configs[n]; !ok {
			panic(errx.New(fmt.Sprintf("db config for <%s> is not set", n)))
		} else {
			gormConfig := &gorm.Config{}

			if c.GormSetters != nil {
				for _, set := range c.GormSetters {
					set(gormConfig)
				}
			}

			var d gorm.Dialector
			d, err := newDialector(c.Dsn)
			if err != nil {
				panic(err)
			}
			db, err = gorm.Open(d, gormConfig)
			if err != nil {
				panic(err)
			}
		}
		return db
	} else {
		return db
	}
}

func newDialector(dsn database.DSN) (gorm.Dialector, error) {
	switch dsn.Type() {
	case database.TypeMysql:
		return mysql.Open(dsn.RemoveSchema()), nil
	case database.TypePostgres:
		return postgres.Open(dsn.RemoveSchema()), nil
	}

	return nil, errx.New(fmt.Sprintf("not support database type <%s>", dsn.Type()))
}

//RefreshDatabase 测试时通过使用本函数来做到不生成实际数据
func RefreshDatabase(f func()) {
	tmp := make(map[string]*gorm.DB)
	for name := range database.Configs {
		if db, ok := dbs[name]; ok {
			tmp[name] = db
		}
		dbs[name] = dbs[name].Begin()
	}
	f()
	for name := range dbs {
		dbs[name].Rollback()
		if db, ok := tmp[name]; ok {
			dbs[name] = db
		} else {
			delete(dbs, name)
		}
	}
}

//Fake 测试时候可以使用gorm的mock工具
func Fake(f func(), GormMockery *gorm.DB) {
	tmp := make(map[string]*gorm.DB)
	for name := range database.Configs {
		if db, ok := dbs[name]; ok {
			tmp[name] = db
		}
		dbs[name] = GormMockery
	}
	f()
	for name := range dbs {
		if db, ok := tmp[name]; ok {
			dbs[name] = db
		} else {
			delete(dbs, name)
		}
	}
}
