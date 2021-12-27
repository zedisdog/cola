package gorm

import (
	"errors"
	"gorm.io/gorm"
)

func NewDBHelper() *DBHelper {
	return &DBHelper{
		DB: DB,
	}
}

type DBHelper struct {
	*gorm.DB
}

//WithTx 没有返回指针是因为一般场景都是用了就丢弃 放到栈上不会给gc压力
func (d *DBHelper) WithTx(tx *gorm.DB) {
	d.DB = tx
}

//Begin not implement
//Deprecated: not implement
func (d DBHelper) Begin() {
	panic(errors.New("not implement"))
}

func (d DBHelper) List(query *gorm.DB, list interface{}, pageSizeOrLimit ...int) (total int, err error) {
	var tmp int64
	err = query.Count(&tmp).Error
	if err != nil {
		return
	}
	total = int(tmp)
	if len(pageSizeOrLimit) > 1 {
		query = query.Offset((pageSizeOrLimit[0] - 1) * pageSizeOrLimit[1]).Limit(pageSizeOrLimit[1])
	} else if len(pageSizeOrLimit) > 0 {
		query = query.Limit(pageSizeOrLimit[0])
	} else {
		//如果没有默认限制100, 防止失误
		query = query.Limit(50)
	}

	err = query.Find(list).Error
	return
}
