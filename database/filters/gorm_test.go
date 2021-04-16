package filters

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestNormal(t *testing.T) {
	var a interface{} = "and"
	if a != "and" {
		t.Fatal("type error")
	}

	b := []string{"and"}
	b[0], b[0] = b[0], b[0]
	fmt.Printf("%+v\n", b)
}

func TestNewFilter(t *testing.T) {
	f := NewFilter("field1", "test")
	if f.field != "field1" ||
		f.value != "test" ||
		f.operator != "=" ||
		f.table != "" ||
		f.on != "" ||
		f.logicOperator != "and" {
		t.Fatalf("%+v", f)
	}

	f = NewFilter("field1", "like", "test")
	if f.field != "field1" ||
		f.value != "test" ||
		f.operator != "like" ||
		f.table != "" ||
		f.on != "" ||
		f.logicOperator != "and" {
		t.Fatal("error2")
	}

	f = NewFilter("field1", "test", "or")
	if f.field != "field1" ||
		f.value != "test" ||
		f.operator != "=" ||
		f.table != "" ||
		f.on != "" ||
		f.logicOperator != "or" {
		t.Fatal("error3")
	}

	f = NewFilter("field1", ">", "test", "or")
	if f.field != "field1" ||
		f.value != "test" ||
		f.operator != ">" ||
		f.table != "" ||
		f.on != "" ||
		f.logicOperator != "or" {
		t.Fatal("error4")
	}

	f = NewFilter("field1", "test", "table", "id=id")
	if f.field != "field1" ||
		f.value != "test" ||
		f.operator != "=" ||
		f.table != "table" ||
		f.on != "id=id" ||
		f.logicOperator != "and" {
		t.Fatalf("%+v", f)
	}

	f = NewFilter("field1", "test", "or", "table", "id=id")
	if f.field != "field1" ||
		f.value != "test" ||
		f.operator != "=" ||
		f.table != "table" ||
		f.on != "id=id" ||
		f.logicOperator != "or" {
		t.Fatalf("%+v", f)
	}

	f = NewFilter("field1", ">", "test", "table", "id=id")
	if f.field != "field1" ||
		f.value != "test" ||
		f.operator != ">" ||
		f.table != "table" ||
		f.on != "id=id" ||
		f.logicOperator != "and" {
		t.Fatalf("%+v", f)
	}

	f = NewFilter("field1", ">", "test", "or", "table", "id=id")
	if f.field != "field1" ||
		f.value != "test" ||
		f.operator != ">" ||
		f.table != "table" ||
		f.on != "id=id" ||
		f.logicOperator != "or" {
		t.Fatalf("%+v", f)
	}
}

func TestNewFilter2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = NewFilter("field1", ">")
}

func TestNewFilter3(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = NewFilter("field1", "a", "table")
}

func TestNewFilter4(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = NewFilter("field1", "a", "or", "table")
}

func TestNewFilter5(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = NewFilter("field1", ">", "a", "table")
}

func TestNewFilter6(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()
	_ = NewFilter("field1", ">", "a", "or", "table")
}

func TestFilters_Run(t *testing.T) {
	total := 100
	db := genDb().Begin()
	defer db.Rollback()
	for i := 0; i < total; i++ {
		test := &Test{
			Name: fmt.Sprintf("record-%d", i),
		}
		err := db.Create(test).Error
		if err != nil {
			t.Fatal(err.Error())
		}
		err = db.Create(&Test2{TestID: test.ID, Name: test.Name}).Error
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	var d Test
	filters := &Filters{
		filters: []*Filter{
			NewFilter("name", "record-1"),
		},
	}
	instance := filters.Apply(db)
	instance.First(&d)
	if d.Name != "record-1" {
		t.Fatal("error1")
	}

	var dd Test
	filters = &Filters{
		filters: []*Filter{
			NewFilter("id", ">", d.ID+1),
		},
	}
	instance = filters.Apply(db)
	instance.First(&dd)
	if dd.Name != "record-3" {
		t.Fatalf("%+v", dd)
	}

	var ddd Test
	filters = &Filters{
		filters: []*Filter{
			NewFilter("name", "like", "%5"),
		},
	}
	instance = filters.Apply(db)
	instance.First(&ddd)
	if ddd.Name != "record-5" {
		t.Fatalf("%+v", ddd)
	}

	var dddd Test2
	filters = &Filters{
		filters: []*Filter{
			NewFilter("tests.name", "like", "%0", "tests", "tests.id=test2.test_id"),
		},
	}
	instance = filters.Apply(db)
	instance.First(&dddd)
	if dddd.Name != "record-0" {
		t.Fatalf("%+v", dddd)
	}
}

func genDb() *gorm.DB {
	dsn := "root:toor@tcp(127.0.0.1:3306)/main?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(Test{}, Test2{})
	if err != nil {
		panic(err)
	}
	return db
}

type Test struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Test2 struct {
	TestID uint
	Name   string
}
