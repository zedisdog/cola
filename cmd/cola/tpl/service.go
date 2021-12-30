package tpl

const ServiceTemp = `package services

import (
	"github.com/zedisdog/cola/tools"
	"gorm.io/gorm"
	"{{.ProjectName}}/internal/{{.ModuleName}}"
	"{{.ProjectName}}/internal/{{.ModuleName}}/dao"
	"{{.ProjectName}}/internal/{{.ModuleName}}/dto"
	"{{.ProjectName}}/internal/{{.ModuleName}}/models"
)

func New{{.Name}}() {{.ModuleName}}.Service {
	return &{{.NameSmallCamel}}{
		{{.NameSmallCamel}}Dao: dao.NewGorm{{.Name}}(),
	}
}

type {{.Name}} struct {
	{{.NameSmallCamel}}Dao {{.ModuleName}}.Dao
}

func (a *{{.NameSmallCamel}}) Create(dto dto.CreateAdmin) (admin models.Admin, err error) {
	err = tools.CopyFields(dto, &admin)
	if err != nil {
		return
	}
	err = a.adminDao.Transaction(func(tx *gorm.DB) (err error) {
		err = a.adminDao.WithTx(tx).Create(&admin)
		if err != nil {
			return
		}
		r, err := a.roleDao.FindByName("admin")
		if err != nil {
			return
		}
		err = a.roleDao.WithTx(tx).RelateRole(admin.ID, r.ID)
		return
	})
	return
}

func (a *adminService) List(dto *dto.ListOptions) (list []models.Admin, total int, err error) {
	return a.adminDao.List(dto)
}

func (a *adminService) Update(dto dto.UpdateAdmin, ID uint64) (admin models.Admin, err error) {
	admin, err = a.adminDao.FindByID(ID)
	if err != nil {
		return
	}
	err = tools.CopyFields(dto, &admin, false)
	if err != nil {
		return
	}
	err = a.adminDao.Update(&admin)
	return
}

func (a *adminService) Delete(ID uint64) error {
	return a.adminDao.DeleteByID(ID)
}

func (a *adminService) FindByID(ID uint64) (admin models.Admin, err error) {
	return a.adminDao.FindByID(ID)
}
`
