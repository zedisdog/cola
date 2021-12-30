package tpl

const ContractTemp = `package {{.PkgName}}

import (
	Gorm "gorm.io/gorm"
	"{{.ProjectName}}/internal/{{.PkgName}}/dto"
	"{{.ProjectName}}/internal/{{.PkgName}}/models"
)

type Dao interface {
	Create(*models.{{.ModelName}}) error
	List(dto *dto.ListOptions) (list []models.{{.ModelName}}, total int, err error)
	Update(*models.{{.ModelName}}) error
	FindByID(ID uint64) (models.{{.ModelName}}, error)
	DeleteByID(ID uint64) error
	Transaction(f func(tx *Gorm.DB) error) error
	WithTx(tx *Gorm.DB) Dao
}

type Service interface {
	Create(dto.Create{{.ModelName}}) (admin models.{{.ModelName}}, err error)
	List(dto *dto.ListOptions) (list []models.{{.ModelName}}, total int, err error)
	Update(*models.{{.ModelName}}) error
	FindByID(ID uint64) (models.{{.ModelName}}, error)
	DeleteByID(ID uint64) error
}
`
