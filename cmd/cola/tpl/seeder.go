package tpl

const SeederTemp = `package seeder

import (
	Gorm "gorm.io/gorm"
)

func {{.SeederName}}(db *Gorm.DB) (err error) {
}
`
