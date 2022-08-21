package main

import (
	"gorm.io/gorm"
	"time"
)

type NotWorn struct {
	Title       string         `json:"title"`
	ID          uint           `gorm:"primaryKey" json:"id"`
	Description string         `json:"description"`
	Condition   string         `json:"condition"`
	Price       uint           `json:"price"`
	CompanyName string         `json:"companyName"`
	TaskForce   string         `json:"taskForce"`
	Location    string         `json:"location"`
	Email       string         `json:"email"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreateAt    time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_At"`
}
