package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgreSQL() (*gorm.DB, error) {
	dsn := "host=localhost user=oscarsolideo password=password dbname=latihan_2 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
