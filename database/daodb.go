package database

import (
	"fmt"
	"github.com/zedisdog/cola/errx"
	"gorm.io/gorm"
)

type DB struct {
	Db    *gorm.DB
	Tx    *gorm.DB
	Conds interface{}
}

func (d DB) get() *gorm.DB {
	if d.Tx != nil {
		return d.Tx
	} else {
		return d.Db
	}
}

func (d *DB) Where(conds ...interface{}) {
	if cond, ok := conds[0].(map[string]interface{}); ok {
		d.Conds = cond
	} else {
		d.Conds = conds
	}
}

func (d *DB) Query() *gorm.DB {
	//every query will be a new query.
	defer func() {
		d.Conds = nil
	}()
	if d.Conds != nil {
		if cMap, ok := d.Conds.(map[string]interface{}); ok {
			return d.get().Where(cMap)
		} else if cSlice, ok := d.Conds.([]interface{}); ok && len(cSlice) > 1 {
			return d.get().Where(cSlice[0], cSlice[1:]...)
		} else {
			panic(errx.New(fmt.Sprintf("unsupported format: %+v", d.Conds)))
		}
	} else {
		return d.get()
	}
}
