package database

import (
	"devbook-api/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.ConnectionDbString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
