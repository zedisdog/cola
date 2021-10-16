package database

import "gorm.io/gorm"

var instances = make(map[string]*gorm.DB)

//Instance return instance of gorm.DB,there are two params:
//	name string instance name
//	Options func(*Options) config function
func Instance(name ...string) *gorm.DB {
	instName := "default"
	if len(name) > 0 {
		instName = name[0]
	}
	return instances[instName]
}

//Init init a default gorm instance for framework
func Init(setOptions ...func(*Options)) {
	db, err := New(setOptions...)
	if err != nil {
		panic(err)
	}
	SetInstance("default", db)
}

func SetInstance(name string, db *gorm.DB) {
	instances[name] = db
}

func ExchangeInstance(name string, db *gorm.DB) (oldInstance *gorm.DB) {
	oldInstance, instances[name] = instances[name], db
	return
}
