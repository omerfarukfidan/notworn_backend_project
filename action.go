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
	obj, err := GetNotWorn(db, id)
	if err != nil {
		log.Println("Error occurred: ", err)
	}
	db.Unscoped().Delete(obj)
	return err
}

func UpdateNotWorn(db *gorm.DB, id any, notworn *NotWorn) (NotWorn, error) {
	obj, err := GetNotWorn(db, id)
	if err != nil {
		log.Println("NotWorn item couldn't find! (UpdateNotWorn)", err)
	}
	err = db.Model(&obj).Updates(map[string]interface{}{"title": &notworn.Title, "description": &notworn.Description, "company_name": &notworn.CompanyName, "condition": &notworn.Condition, "price": &notworn.Price, "location": &notworn.Location}).Error
	if err != nil {
		log.Println("Error occurred (UpdateNotWorn):  ", err)
	}
	/*
		err = db.Model(&obj).Updates(&updatedFields).Error
		if err != nil {
			log.Println("Error occurred (UpdateNotWorn):  ", err)
		}

	*/
	//"title": notworn.Title, "description": notworn.Description, "company_name": notworn.CompanyName, "condition": notworn.Condition, "price": notworn.Price, "location": notworn.Location, "deleted_at": notworn.DeletedAt, "updated_at": notworn.UpdatedAt
	return obj, err
}
