package main

import (
	"gorm.io/gorm"
	"log"
)

func CreateNotWorn(db *gorm.DB, notworn *NotWorn) error {
	err := db.Create(&notworn).Error
	if err != nil {
		log.Println("Error occurred (CreateNotWorn):", err)
	}
	return err
}
