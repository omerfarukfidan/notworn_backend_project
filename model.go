package main

import (
	"gorm.io/gorm"
	"mime/multipart"
	"time"
)

type NotWorn struct {
	DeletedAt   gorm.DeletedAt `gorm:"index" form:"deleted_at"` // 32 Bytes
	CreateAt    time.Time      `form:"created_at"`              // 24 Bytes
	UpdatedAt   time.Time      `form:"updated_at"`              // 24 Bytes
	Title       string         `form:"title"`                   // 16 Bytes
	Description string         `form:"description"`             // 16 Bytes
	Condition   string         `form:"condition"`               // 16 Bytes
	CompanyName string         `form:"companyName"`             // 16 Bytes
	Location    string         `form:"location"`                // 16 Bytes
	FileName    string         `form:"filename"`                // 16 Bytes
	ImagePath   string         `form:"-"`                       // 16 Bytes
	ID          uint           `gorm:"primaryKey" form:"id"`    // 8 Bytes
	Price       float64        `form:"price"`                   // 8 Bytes

}
type file struct {
	ID     uint                  `form:"id"`                        // 8 Bytes
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"` // 8 Bytes
}
