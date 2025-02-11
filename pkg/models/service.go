package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name     string
	NodeType string
	Address  string
	Port     int
	Status   string
}
