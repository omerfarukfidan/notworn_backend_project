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
	Location    string         `form:"location"`
	FileName    string         `form:"filename"`
	ImagePath   string         `form:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" form:"deleted_at"`
	CreateAt    time.Time      `form:"created_at"`
	UpdatedAt   time.Time      `form:"updated_at"`
}
type file struct {
	ID     uint                  `form:"id"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}
