package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name     string `gorm:"unique"`
	NodeType string
	Address  string
	Port     string
	Status   string
}
