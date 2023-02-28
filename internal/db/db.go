package db

import (
	"GinProject/internal/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbSession *gorm.DB
)

func Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	ds, err := gorm.Open(postgres.New(postgres.Config{
		DSN: os.Getenv("DB_CONFIG"),
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	ds.AutoMigrate(&model.Users{})

	dbSession = ds
}

func GetDBSession() *gorm.DB {
	return dbSession
}
