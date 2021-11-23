package seeder

import "gorm.io/gorm"

var s = make(seeders, 0)

type seeders []func(db *gorm.DB) error

func Seed(db *gorm.DB) {
	for _, seed := range s {
		err := seed(db)
		if err != nil {
			panic(err)
		}
	}
	s = nil // 释放内存
	return
}

func Add(seeders ...func(db *gorm.DB) error) {
	s.add(seeders...)
}
func (s *seeders) add(seeders ...func(db *gorm.DB) error) {
	*s = append(*s, seeders...)
}
