package database

import (
	"fmt"
	"github.com/zedisdog/cola/errx"
	"gorm.io/gorm"
)

type DB struct {
	Db       *gorm.DB
	Tx       *gorm.DB
	Conds    []interface{}
	Joins    []string
	Preloads []string
	Offset   *int
	Limit    int
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

func (d *DB) Join(conds ...string) {
	d.Joins = append(d.Joins, conds...)
}

func (d *DB) Preload(relates ...string) {
	d.Preloads = append(d.Preloads, relates...)
}

func (d *DB) setOffset(offset int) {
	d.Offset = &offset
}

func (d *DB) setLimit(limit int) {
	d.Limit = limit
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
		d.Preloads = nil
		d.Offset = nil
		d.Limit = 0
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

	for _, p := range d.Preloads {
		query = query.Preload(p)
	}

	if d.Offset != nil {
		query = query.Offset(*d.Offset)
	}

	if d.Limit != 0 {
		query = query.Limit(d.Limit)
	}

	return query
}
