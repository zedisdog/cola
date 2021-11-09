package database

import (
	"fmt"
	"github.com/zedisdog/cola/errx"
	"gorm.io/gorm"
)

type DB struct {
	Db    *gorm.DB
	Tx    *gorm.DB
	Conds []interface{}
	Joins []string
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
	if d.Conds == nil {
		d.Conds = make([]interface{}, 0, 5)
	}
	if cond, ok := conds[0].(map[string]interface{}); ok {
		d.Conds = append(d.Conds, cond)
	} else {
		d.Conds = append(d.Conds, conds)
	}
}

func (d *DB) Join(conds string) {
	d.Joins = append(d.Joins, conds)
}

//Deprecated: use Builder instead
func (d *DB) Query() *gorm.DB {
	return d.Builder()
}

//Builder return an instance of gorm.DB, which with simple query conditions set by DB.Where.
func (d *DB) Builder() *gorm.DB {
	defer func() {
		d.Conds = nil
		d.Joins = nil
	}()
	query := d.get()
	for _, c := range d.Conds {
		if cMap, ok := c.(map[string]interface{}); ok {
			query = query.Where(cMap)
		} else if cSlice, ok := c.([]interface{}); ok && len(cSlice) > 1 {
			query = query.Where(cSlice[0], cSlice[1:]...)
		} else {
			panic(errx.New(fmt.Sprintf("unsupported format: %+v", d.Conds)))
		}
	}

	for _, c := range d.Joins {
		query = query.Joins(c)
	}

	return query
}
