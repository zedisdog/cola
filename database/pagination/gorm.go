package pagination

import (
	"gorm.io/gorm"
	"reflect"
)

func NewGormPaginator(db *gorm.DB, conditionsAndBinds ...interface{}) *GormPaginator {
	gp := &GormPaginator{
		db: db,
	}
	if len(conditionsAndBinds) >= 2 {
		gp.conditions = conditionsAndBinds[0].(string)
		gp.binds = conditionsAndBinds[1:]
	}
	return gp
}

type GormPaginator struct {
	db         *gorm.DB
	conditions string
	binds      []interface{}
}

func (g GormPaginator) Page(list interface{}, currentPage int, pageSize int) (total int, err error) {
	var count int64
	err = g.db.Model(g.newStructWithSlice(list)).Where(g.conditions, g.binds...).Count(&count).Error
	if err != nil {
		return
	}
	total = int(count)
	err = g.db.Where(g.conditions, g.binds...).Offset(pageSize * (currentPage - 1)).Limit(pageSize).Find(list).Error
	return
}

func (g GormPaginator) newStructWithSlice(ptr interface{}) interface{} {
	t := reflect.ValueOf(ptr).Elem().Type().Elem()
	return reflect.New(t).Interface()
}
