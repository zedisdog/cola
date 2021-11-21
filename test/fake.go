package test

import (
	"github.com/zedisdog/cola/database"
	"gorm.io/gorm"
	"sync"
)

var fakeDB *gorm.DB
var unDo = make([]func(), 0, 1)

//FakeDB use a fake DBHelper instance.
func FakeDB() {
	var once sync.Once
	once.Do(func() {
		var err error
		fakeDB, err = database.NewMocker()
		if err != nil {
			panic(err)
		}
	})
	oldInstance := database.ExchangeInstance("default", fakeDB)
	unDo = append(unDo, func() {
		database.ExchangeInstance("default", oldInstance)
	})
}

func Fake(test func(), fakes ...func()) {
	// enable fakes
	for _, f := range fakes {
		f()
	}
	defer func() { // clear fakes
		for _, u := range unDo {
			u()
		}
	}()
	test()
}
