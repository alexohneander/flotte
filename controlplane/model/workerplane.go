package model

import (
	"gorm.io/gorm"
)

type WorkerPlane struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	Name    string `gorm:"unique"`
	IP      string
	APIPort string
	APIKey  string
	Status  string
}
