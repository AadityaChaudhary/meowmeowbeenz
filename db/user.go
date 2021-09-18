package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Score   float32
	Picture string
	Name    string
}

func NewUser(Picture, Name string) *User {
	if Name == SYS_NAME {
		return &User{}
	}
	return &User{
		Picture: Picture,
		Name:    Name,
		Score:   3.0,
	}
}
