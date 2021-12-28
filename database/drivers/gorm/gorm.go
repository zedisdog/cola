package gorm

import (
	"errors"
	"github.com/zedisdog/cola/tools"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
)

var DB *gorm.DB

func InitDB(dsn string, setters ...func(*gorm.Config)) (db *gorm.DB, err error) {
	if DB == nil {
		config := &gorm.Config{}
		for _, set := range setters {
			set(config)
		}

		var d gorm.Dialector
		d, err = newDialector(tools.EncodeQuery(dsn))
		if err != nil {
			panic(err)
		}
		DB, err = gorm.Open(d, config)
	}

	db = DB

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
	tmp := DB
	DB = DB.Begin()
	f()
	DB.Rollback()
	DB = tmp
}

//Fake 测试时候可以使用gorm的mock工具
func Fake(f func(), GormMockery *gorm.DB) {
	tmp := DB
	DB = GormMockery
	f()
	DB = tmp
}
