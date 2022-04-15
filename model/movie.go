package model

import (
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ImageURL    string         `json:"image_url"`
	Director    string         `json:"director"`
	CreateAt    time.Time      `json:"created"`
	UpdateAt    time.Time      `json:"Updated"`
	DeleteAt    gorm.DeletedAt `gorm:"index" json:"deleted"`
}
