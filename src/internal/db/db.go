package db

import (
	"fmt"

	"se-api/src/internal/config"

	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var m *migrate.Migrate
	var err error

	h := config.AppConfig.DB_HOST
	d := config.AppConfig.DB_DATABASE
	u := config.AppConfig.DB_USERNAME
	p := config.AppConfig.DB_PASSWORD
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=True&loc=Local", u, p, h, d)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error initializing db: %v", err)
		return err
	}

	// migration
	m, err = migrate.New(config.AppConfig.MIGRATION_FOLDER_PATH, "mysql://"+dsn)
	if err != nil {
		log.Printf("Error initializing migration: %v", err)
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Printf("Error migration: %v", err)
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
