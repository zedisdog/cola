package filters

import (
	"errors"
	"fmt"
	"github.com/zedisdog/cola/tools"
	"gorm.io/gorm"
)

// NewFilter new a Filter instance
//  usage:
//    NewFilter(field,[operator],value,[logicOperator,relateTable,relateTableCondition])
func NewFilter(params ...interface{}) *Filter {
	f := &Filter{
		field: params[0].(string),
	}

	switch len(params) {
	case 6:
		f.field = params[0].(string)
		f.operator = params[1].(string)
		f.value = params[2]
		f.logicOperator = params[3].(string)
		f.table = params[4].(string)
		f.on = params[5].(string)
	case 5:
		if !tools.InArray(params[1], []interface{}{">", "<", "=", "like", "!="}) {
			f.field = params[0].(string)
			f.operator = "="
			f.value = params[1]
			f.logicOperator = params[2].(string)
			f.table = params[3].(string)
			f.on = params[4].(string)
		} else if !tools.InArray(params[3], []interface{}{"and", "or"}) {
			f.field = params[0].(string)
			f.operator = params[1].(string)
			f.value = params[2]
			f.logicOperator = "and"
			f.table = params[3].(string)
			f.on = params[4].(string)
		} else {
			panic(errors.New("relate table should with condition"))
		}
	case 4:
		if tools.InArray(params[1], []interface{}{">", "<", "=", "like", "!="}) &&
			tools.InArray(params[3], []interface{}{"and", "or"}) {
			f.field = params[0].(string)
			f.operator = params[1].(string)
			f.value = params[2]
			f.logicOperator = params[3].(string)
		} else if !tools.InArray(params[1], []interface{}{">", "<", "=", "like", "!="}) &&
			!tools.InArray(params[2], []interface{}{"and", "or"}) {
			f.field = params[0].(string)
			f.operator = "="
			f.value = params[1]
			f.logicOperator = "and"
			f.table = params[2].(string)
			f.on = params[3].(string)
		} else {
			panic(errors.New("relate table should with condition1"))
		}
	case 3:
		if tools.InArray(params[1], []interface{}{">", "<", "=", "like", "!="}) {
			f.field = params[0].(string)
			f.operator = params[1].(string)
			f.value = params[2]
			f.logicOperator = "and"
		} else if tools.InArray(params[2], []interface{}{"and", "or"}) {
			f.field = params[0].(string)
			f.operator = "="
			f.value = params[1]
			f.logicOperator = params[2].(string)
		} else {
			panic(errors.New("relate table should with condition2"))
		}
	case 2:
		if !tools.InArray(params[1], []interface{}{">", "<", "=", "like", "!="}) {
			f.field = params[0].(string)
			f.operator = "="
			f.value = params[1]
			f.logicOperator = "and"
		} else {
			panic(errors.New("relate table should with condition3"))
		}
	}

	return f
}

type Filter struct {
	table         string
	on            string
	field         string
	operator      string
	value         interface{}
	logicOperator string
}

func (f Filter) Relate(db *gorm.DB) {
	if f.table != "" && f.on != "" {
		db.Joins(fmt.Sprintf("left join %s on %s", f.table, f.on))
	}
}

func (f Filter) Query() string {
	return fmt.Sprintf("%s %s ?", f.field, f.operator)
}

func (f Filter) Value() interface{} {
	return f.value
}

type Filters struct {
	filters []*Filter
}

func (f Filters) Apply(db *gorm.DB) *gorm.DB {
	sql := ""
	binds := make([]interface{}, 0, len(f.filters))

	// 第一个必须是and 或者 全部都是 or
	for index, filter := range f.filters {
		if filter.logicOperator == "and" {
			f.filters[index], f.filters[0] = f.filters[0], f.filters[index]
			break
		}
	}

	for _, filter := range f.filters {
		if sql == "" {
			sql += fmt.Sprintf("%s %s ?", filter.field, filter.operator)
		} else {
			sql += fmt.Sprintf(" %s %s %s ?", filter.logicOperator, filter.field, filter.operator)
		}

		binds = append(binds, filter.value)

		if filter.table != "" && filter.on != "" {
			db = db.Joins(fmt.Sprintf("LEFT JOIN %s ON %s", filter.table, filter.on))
		}
	}
	return db.Where(sql, binds...)
}
