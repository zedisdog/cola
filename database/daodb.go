package database

import "gorm.io/gorm"

type DB struct {
	Db    *gorm.DB
	Tx    *gorm.DB
	Conds map[string]interface{}
}

func (d DB) get() *gorm.DB {
	if d.Tx != nil {
		return d.Tx
	} else {
		return d.Db
	}
}

func (d *DB) Query() *gorm.DB {
	defer func() {
		d.Conds = nil
	}()
	if d.Conds != nil {
		return d.get().Where(d.Conds)
	} else {
		return d.get()
	}
}
