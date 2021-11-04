package database

import (
	"github.com/zedisdog/cola/tools"
	"gorm.io/gorm"
	"time"
)

//Model is gorm model helper struct with snowflakeID and time.Time.
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

//ModelWithUnixTimeStamp is gorm model helper struct with snowflakeID and int64.
type ModelWithUnixTimestamp struct {
	ID        uint64 `gorm:"primary_key" json:"id,string"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (m *ModelWithUnixTimestamp) BeforeCreate(tx *gorm.DB) error {
	if m.ID == 0 {
		id, err := tools.GetSnow().NextID()
		if err != nil {
			return err
		}
		m.ID = id
	}
	now := time.Now().Unix()
	if m.CreatedAt == 0 {
		m.CreatedAt = now
	}
	if m.UpdatedAt == 0 {
		m.UpdatedAt = now
	}
	return nil
}

func (m *ModelWithUnixTimestamp) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now().Unix()
	if m.UpdatedAt == 0 {
		m.UpdatedAt = now
	}
	return nil
}

func (m *ModelWithUnixTimestamp) CreatedAtTime() time.Time {
	return time.Unix(m.CreatedAt, 0)
}

func (m *ModelWithUnixTimestamp) UpdatedAtTime() time.Time {
	return time.Unix(m.UpdatedAt, 0)
}
