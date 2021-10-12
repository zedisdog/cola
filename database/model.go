package database

import (
	"github.com/zedisdog/cola/tools"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint64     `gorm:"primary_key" json:"id,string"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	if m.ID == 0 {
		id, err := tools.GetSnow().NextID()
		if err != nil {
			return err
		}
		m.ID = id
	}
	return nil
}
