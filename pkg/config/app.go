package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load("D:\\gousermodules\\.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		host, user, password, dbname, port)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
