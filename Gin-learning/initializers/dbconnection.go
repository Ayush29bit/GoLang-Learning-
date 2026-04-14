package initializers

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionToDB() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	dsn := os.Getenv("DB_URL")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	_ = DB
}
