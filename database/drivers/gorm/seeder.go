package gorm

import "gorm.io/gorm"

type SeedFunc func(db *gorm.DB) error

var seeders = make([]SeedFunc, 0)

func Add(f ...SeedFunc) {
	seeders = append(seeders, f...)
}

func Seed(db *gorm.DB) (err error) {
	for _, seeder := range seeders {
		err = seeder(db)
		if err != nil {
			return
		}
	}

	return
}
