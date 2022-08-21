package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDataBase() (*gorm.DB, error) {
	var notworn NotWorn
	var db *gorm.DB

	postgresqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_TZ"))
	db, err := gorm.Open(postgres.Open(postgresqlInfo))

	if err != nil {
		log.Fatalln("Failed to connect database!")
	}

	err = db.AutoMigrate(&notworn)
	if err != nil {
		log.Fatalln("error was:", err)
	}

	return db, err
}
