package model

import (
	"gorm.io/gorm"
)

type ControlPlane struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	Name       string `gorm:"unique"`
	IP         string `gorm:"unique"`
	APIPort    string
	APIKey     string
	Status     string
	NodeCount  int
	LastUpdate string `gorm:"autoUpdateTime"`
}
