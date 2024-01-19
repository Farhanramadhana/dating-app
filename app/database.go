package app

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	host := os.Getenv("postgre_host")
	user := os.Getenv("postgre_user")
	password := os.Getenv("postgre_password")
	dbname := os.Getenv("postgre_dbname")
	port := os.Getenv("postgre_port")
	sslmode := os.Getenv("postgre_sslmode")
	timezone := os.Getenv("postgre_timezone")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
