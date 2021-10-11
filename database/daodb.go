package database

import "gorm.io/gorm"

type DB struct {
	Db *gorm.DB
	tx *gorm.DB
}

func (d DB) Get() *gorm.DB {
	if d.tx != nil {
		return d.tx
	} else {
		return d.Db
	}
}

func (d DB) SetDb(db *gorm.DB) {
	d.Db = db
}

func (d DB) SetTx(tx *gorm.DB) {
	d.tx = tx
}
