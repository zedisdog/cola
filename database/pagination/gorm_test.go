package pagination

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Test struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func TestGorm_Page(t *testing.T) {
	currentPage := 1
	pageSize := 10
	total := 100
	db := genDb().Begin()
	defer db.Rollback()
	for i := 0; i < total; i++ {
		err := db.Create(&Test{
			Name: fmt.Sprintf("record-%d", i),
		}).Error
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	driver := NewGormPaginator(db, "")
	var list []Test
	actuallyTotal, err := driver.Page(&list, currentPage, pageSize)
	if err != nil {
		t.Fatal(err.Error())
	}
	if total != actuallyTotal {
		t.Fatalf("should be %d, actually %d", total, actuallyTotal)
	}
	if len(list) != pageSize {
		t.Fatalf("list len should be %d, actually %d", pageSize, len(list))
	}
}

func TestGorm_newStructWithSlice(t *testing.T) {
	driver := NewGormPaginator(nil, "")
	var a []Test
	s := driver.newStructWithSlice(&a)
	if _, ok := s.(*Test); !ok {
		t.Fatal("get type error")
	}
}

func genDb() *gorm.DB {
	dsn := "root:toor@tcp(127.0.0.1:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(Test{})
	if err != nil {
		panic(err)
	}
	return db
}
