package tpl

var DaoTemp = `package dao

import (
	"github.com/zedisdog/cola/database/drivers/gorm"
	"github.com/zedisdog/cola/errx"
	Gorm "gorm.io/gorm"
	"{{.PkgName}}/internal/{{.ModuleName}}"
	"{{.PkgName}}/internal/{{.ModuleName}}/dto"
	"{{.PkgName}}/internal/{{.ModuleName}}/model"
)

func NewGorm{{.Entity}}Dao() {{.ModuleName}}.Dao {
	return &gorm{{.Entity}}Dao{
		db: gorm.NewDBHelper(),
	}
}

type gorm{{.Entity}}Dao struct {
	db *gorm.DBHelper
}

func (g gorm{{.Entity}}Dao) WithTx(tx *Gorm.DB) {{.ModuleName}}.Dao {
	m.db.WithTx(tx)
	return &m
}

func (g *gorm{{.Entity}}Dao) Transaction(f func(tx *Gorm.DB) error) error {
	return m.db.Transaction(f)
}

func (g *gorm{{.Entity}}Dao) Create({{.EntitySmallCamel}} *models.{{.Entity}}) (err error) {
	return errx.Wrap(m.db.Create({{.EntitySmallCamel}}).Error, "create {{.EntitySmallCamel}} failed")
}

func (g *gorm{{.Entity}}Dao) List(dto *dto.ListOptions) (list []models.{{.Entity}}, total int, err error) {
	query := m.db.Model(&models.{{.Entity}}{}).Order("id DESC")

	for key, value := range dto.Filters {
		switch key {
		default:
		}
	}

	if dto.Page != 0 {
		if dto.Size == 0 {
			dto.Size = 15
		}
		total, err = m.db.List(query, &list, dto.Page, dto.Size)
	} else if dto.Limit != 0 {
		_, err = m.db.List(query, &list, dto.Limit)
	}

	if err != nil {
		err = errx.Wrap(err, "list {{.EntitySmallCamel}} failed")
	}

	return
}

func (g *gorm{{.Entity}}Dao) Update({{.EntitySmallCamel}} *models.{{.Entity}}) error {
	return m.db.Select("*").Updates({{.EntitySmallCamel}}).Error
}

func (g *gorm{{.Entity}}Dao) DeleteByID(ID uint64) error {
	return m.db.Delete(&models.{{.Entity}}{}, ID).Error
}

func (g *gorm{{.Entity}}Dao) FindByID(ID uint64) ({{.EntitySmallCamel}} models.{{.Entity}}, err error) {
	err = m.db.Find(&{{.EntitySmallCamel}}, ID).Error
	return
}
`
