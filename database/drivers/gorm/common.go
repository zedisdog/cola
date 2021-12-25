package gorm

import (
	"github.com/zedisdog/cola/tools"
	"gorm.io/gorm"
	"time"
)

type CommonField struct {
	ID        uint64 `json:"id,string" gorm:"primary"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (c *CommonField) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == 0 {
		c.ID, err = tools.GetSnow().NextID()
		if err != nil {
			return
		}
	}

	now := time.Now()
	if c.CreatedAt == 0 {
		c.CreatedAt = now.Unix()
	}

	if c.UpdatedAt == 0 {
		c.UpdatedAt = now.Unix()
	}

	return
}

func (c *CommonField) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()

	c.UpdatedAt = now.Unix()

	return
}
