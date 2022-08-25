package main

import (
	"fmt"
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
		log.Println("Error occurred ( GetNotWorn): ", err)
	}

	return notworn, err
}

func WriteFileName(db *gorm.DB, obj *NotWorn, userObj *file) error {
	stringID := fmt.Sprintf("%d", obj.ID)
	err := db.Model(&obj).Update("file_name", stringID+userObj.Avatar.Filename).Error
	if err != nil {
		log.Println("Error occurred (WriteFileName) : ", err)
	}
	return err
}

func ListAllNotWorn(db *gorm.DB) ([]NotWorn, error) {
	var notworn []NotWorn
	err := db.Find(&notworn).Error

	if err != nil {
		log.Println("Error occurred (ListAllNotWorn): ", err)
	}

	return notworn, err
}

func HardDeleteNotWorn(db *gorm.DB, id any) error {
	var obj NotWorn
	err := db.Unscoped().Delete(&obj, id).Error
	if err != nil {
		log.Println("Error occurred: ", err)
	}

	return err
}

func DeleteNotWorn(db *gorm.DB, id any) error {
	var obj NotWorn
	err := db.Delete(&obj, id).Error
	if err != nil {
		log.Println("Error occurred: ", err)
	}
	return err
}
