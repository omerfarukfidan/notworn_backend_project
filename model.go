package main

import (
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type NotWorn struct {
	Title       string         `form:"title"`
	ID          uint           `gorm:"primaryKey" form:"id"`
	Description string         `form:"description"`
	Condition   string         `form:"condition"`
	Price       uint           `form:"price"`
	CompanyName string         `form:"companyName"`
	TaskForce   string         `form:"taskForce"`
	Location    string         `form:"location"`
	Email       string         `form:"email"`
	FileName    string         `form:"filename"`
	DeletedAt   gorm.DeletedAt `gorm:"index" form:"deleted_at"`
	CreateAt    time.Time      `form:"created_at"`
	UpdatedAt   time.Time      `form:"updated_At"`
}
type file struct {
	ID     uint                  `form:"id"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}
