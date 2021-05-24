package seeder

import "gorm.io/gorm"

var s = make(seeders, 0)

type seeders []func(db *gorm.DB) error

func Seed(db *gorm.DB) {
	s.seed(db)
}
func (s seeders) seed(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, seed := range s {
			err := seed(tx)
			if err != nil {
				return err
			}
		}
		s = nil // 释放内存
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func Add(seeders ...func(db *gorm.DB) error) {
	s.add(seeders...)
}
func (s *seeders) add(seeders ...func(db *gorm.DB) error) {
	*s = append(*s, seeders...)
}
