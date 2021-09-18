package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var SystemID uint

const SYS_NAME string = "meow meow beenz"

func Start(connString string) error {
	fmt.Println("connecting to db")
	var err error
	DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return err
	}
	DB.AutoMigrate(&User{}, &Rating{})
	var u User
	DB.FirstOrCreate(&u, User{Name: SYS_NAME})
	SystemID = u.ID
	return nil
}
