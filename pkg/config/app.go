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

type Env struct {
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabasePort     string
	JWTSecret        string
}

func LoadConfig() (*Env, error) {
	err := godotenv.Load("D:\\gousermodules\\.env") // chnage path of .env
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return nil, err
	}
	return &Env{
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		JWTSecret:        os.Getenv("SECRET_KEY"),
	}, nil
}

func Connect() {
	env, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		env.DatabaseHost, env.DatabaseUser, env.DatabasePassword, env.DatabaseName, env.DatabasePort)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
