package tpl

const ModelTemp = `package models

import "github.com/zedisdog/cola/database/drivers/gorm"

type {{.ModelName}} struct {
	gorm.CommonField
}
`
