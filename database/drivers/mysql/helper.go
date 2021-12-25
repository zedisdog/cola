package mysql

import (
	"context"
	"database/sql"
	"errors"
)

func NewDBHelper() *DBHelper {
	return &DBHelper{
		DB: DB,
	}
}

type DBHelper struct {
	*sql.DB
	*sql.Tx
}

//WithTx 没有返回指针是因为一般场景都是用了就丢弃 放到栈上不会给gc压力
func (d DBHelper) WithTx(tx *sql.Tx) DBHelper {
	d.Tx = tx
	return d
}

func (d *DBHelper) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	//todo: wrap
	return nil, errors.New("not implement")
}

func (d *DBHelper) Begin() (*sql.Tx, error) {
	//todo: wrap
	return nil, errors.New("not implement")
}

func (d *DBHelper) Transaction(f func(tx *sql.Tx) error) error {
	if d.Tx != nil {
		return f(d.Tx)
	} else {
		tx, err := d.DB.Begin()
		if err != nil {
			return err
		}
		err = f(tx)
		if err != nil {
			e := tx.Rollback()
			if e != nil {
				return e
			} else {
				return err
			}
		}
		return tx.Commit()
	}
}

func (d *DBHelper) Prepare(query string) (*sql.Stmt, error) {
	if d.Tx != nil {
		return d.Tx.Prepare(query)
	} else {
		return d.DB.Prepare(query)
	}
}

func (d *DBHelper) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if d.Tx != nil {
		return d.Tx.PrepareContext(ctx, query)
	} else {
		return d.DB.PrepareContext(ctx, query)
	}
}

func (d *DBHelper) Exec(query string, args ...interface{}) (sql.Result, error) {
	if d.Tx != nil {
		return d.Tx.Exec(query, args...)
	} else {
		return d.DB.Exec(query, args...)
	}
}

func (d *DBHelper) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if d.Tx != nil {
		return d.Tx.ExecContext(ctx, query, args...)
	} else {
		return d.DB.ExecContext(ctx, query, args...)
	}
}

func (d *DBHelper) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if d.Tx != nil {
		return d.Tx.Query(query, args...)
	} else {
		return d.DB.Query(query, args...)
	}
}

func (d *DBHelper) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if d.Tx != nil {
		return d.Tx.QueryContext(ctx, query, args...)
	} else {
		return d.DB.QueryContext(ctx, query, args)
	}
}

func (d *DBHelper) QueryRow(query string, args ...interface{}) *sql.Row {
	if d.Tx != nil {
		return d.Tx.QueryRow(query, args...)
	} else {
		return d.DB.QueryRow(query, args...)
	}
}

func (d *DBHelper) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if d.Tx != nil {
		return d.Tx.QueryRowContext(ctx, query, args...)
	} else {
		return d.DB.QueryRowContext(ctx, query, args...)
	}
}
