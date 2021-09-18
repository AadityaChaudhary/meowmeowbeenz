package main

import (
	"log"
	"os"
	"socialcredit/db"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("starting")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Start(os.Getenv("CRDB"))
	if err != nil {
		log.Fatal(err)
	}
	//	user := &models.User{Name: "aadi", Picture: "", Score: 3.0}
	var u db.User
	db.DB.First(&u, 694284238361372433)
	log.Println("name", u.Name)
	//db.Create(user)
	//log.Println("id", user.ID)
}
