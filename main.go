package main

import (
	"fmt"
	"log"
	"os"
	"socialcredit/db"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("starting")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Start(os.Getenv("CRDB"))
	if err != nil {
		log.Fatal(err)
	}
}
