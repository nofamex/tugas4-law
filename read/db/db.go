package db

import (
	"fmt"
	"log"
	"read/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.ApiConfig.DBHost,
		config.ApiConfig.DBUser,
		config.ApiConfig.DBPassword,
		config.ApiConfig.DBName,
		config.ApiConfig.DBPort,
	)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	MigrateDatabase(DB)
	return DB, err
}

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&User{},
	)
	if err != nil {
		log.Fatal("failed migrating database: ", err.Error())
	}
}