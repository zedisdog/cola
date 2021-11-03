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

//Where set simple conditions supported by gorm.
//   example:
//     DB.Where("id = ?", 1) => gorm.DB.Where("id = ?", 1)
//     DB.Where(map[string]interface{"id": 1}) => gorm.DB.Where(map[string]interface{"id": 1})
func (d *DB) Where(conds ...interface{}) {
	if cond, ok := conds[0].(map[string]interface{}); ok {
		d.Conds = cond
	} else {
		d.Conds = conds
	}
}

//Deprecated: use Builder instead
func (d *DB) Query() *gorm.DB {
	return d.Builder()
}

//Builder return an instance of gorm.DB, which with simple query conditions set by DB.Where.
func (d *DB) Builder() *gorm.DB {
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
