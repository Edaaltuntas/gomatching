package models

import "gorm.io/gorm"

// A basic model to be able to store registers
type Person struct {
	gorm.Model
	Name          string `gorm:"type: text;not null"`
	Category      string `gorm:"type: text;not null"`
	City          string `gorm:"type: text;not null"`
	MatchingScore uint8  `gorm:"-"`
}
