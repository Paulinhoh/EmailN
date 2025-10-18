package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=lmz1507 dbname=emailn port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connect to database")
	}
	return db
}
