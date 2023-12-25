package db

import (
	"fmt"
	db_models "internal-api/db/models"
	"os"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_Connection *gorm.DB

func ConnectToDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Yekaterinburg",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.Errorf("Failed to connect database: %v", err)
	}

	DB_Connection = db

	return nil
}

func Migrate() error {
	return DB_Connection.AutoMigrate(
		&db_models.User{},
		&db_models.Video{},
		&db_models.MediaSource{},
		&db_models.Comment{},
		&db_models.Like{},
	)
}
