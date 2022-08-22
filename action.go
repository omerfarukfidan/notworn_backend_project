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

func GetNotWorn(db *gorm.DB, id any) (NotWorn, error) {
	var notworn NotWorn
	err := db.First(&notworn, id).Error

	if err != nil {
		log.Println("Error occurred (GetTodo): ", err)
	}

	return notworn, err
}

func WriteFileName(db *gorm.DB, obj *NotWorn, userObj *file) error {
	err := db.Model(&obj).Update("file_name", userObj.Avatar.Filename).Error
	if err != nil {
		log.Println("Error occurred : ", err)
	}
	return err
}
