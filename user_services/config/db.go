package config

import (
	"fmt"
	"os"

	"user-services/models"

	spannergorm "github.com/googleapis/go-gorm-spanner"
	"gorm.io/gorm"
)

func getDatabaseString() string {
	return fmt.Sprintf("projects/%s/instances/%s/databases/%s", os.Getenv("PROJECT_ID"), os.Getenv("INSTANCE_ID"), os.Getenv("DATABASE_ID"))
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(spannergorm.New(spannergorm.Config{DriverName: "spanner", DSN: getDatabaseString()}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{})
	return db
}
