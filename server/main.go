package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/adelapazborrero/gophisher/config"
	"github.com/adelapazborrero/gophisher/mail"
	"github.com/adelapazborrero/gophisher/migration"
)

func main() {
	db, err := config.InitDB([]string{
		migration.CreateMailTable,
		migration.CreateMailTable,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reader := mail.NewCSVReader()
	emails, err := reader.GetEmailsFromFile("./mail.txt")
	if err != nil {
		log.Fatal("Could not read emails from file. ", err)
	}

	fmt.Println(emails)
	// server := config.InitHTTP()
	// server.Run()

}
