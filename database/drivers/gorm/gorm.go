package gorm

import (
	"errors"
	"github.com/zedisdog/cola/tools"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
)

var db *gorm.DB

func Instance() *gorm.DB {
	return db
}

func Init(dsn string, setters ...func(*gorm.Config)) (err error) {
	if db == nil {
		config := &gorm.Config{}
		for _, set := range setters {
			set(config)
		}

		var d gorm.Dialector
		d, err = newDialector(tools.EncodeQuery(dsn))
		if err != nil {
			panic(err)
		}
		db, err = gorm.Open(d, config)
	}

	return
}

func newDialector(dsn string) (gorm.Dialector, error) {
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

//RefreshDatabase 测试时通过使用本函数来做到不生成实际数据
func RefreshDatabase(f func()) {
	tmp := db
	db = db.Begin()
	f()
	db.Rollback()
	db = tmp
}

//Fake 测试时候可以使用gorm的mock工具
func Fake(f func(), GormMockery *gorm.DB) {
	tmp := db
	db = GormMockery
	f()
	db = tmp
}
