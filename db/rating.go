package db

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Subject uint
	Object  uint
	Score   int
}
