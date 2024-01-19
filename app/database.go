package app

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	host := os.Getenv("postgres_host")
	user := os.Getenv("postgres_user")
	password := os.Getenv("postgres_password")
	dbname := os.Getenv("postgres_dbname")
	port := os.Getenv("postgres_port")
	sslmode := os.Getenv("postgres_sslmode")
	timezone := os.Getenv("postgres_timezone")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
