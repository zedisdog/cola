package pagination

import (
	"gorm.io/gorm"
	"reflect"
)

func NewGormPaginator(db *gorm.DB, conditions string, binds ...interface{}) Paginator {
	return &Gorm{
		db,
		conditions,
		binds,
	}
}

type Gorm struct {
	db         *gorm.DB
	conditions string
	binds      []interface{}
}

func (g Gorm) Page(list interface{}, currentPage int, pageSize int) (total int, err error) {
	var count int64
	err = g.db.Model(g.newStructWithSlice(list)).Where(g.conditions, g.binds...).Count(&count).Error
	if err != nil {
		return
	}
	total = int(count)
	err = g.db.Where(g.conditions, g.binds...).Offset(pageSize * (currentPage - 1)).Limit(pageSize).Find(list).Error
	return
}

func (g Gorm) newStructWithSlice(ptr interface{}) interface{} {
	t := reflect.ValueOf(ptr).Elem().Type().Elem()
	return reflect.New(t).Interface()
}
