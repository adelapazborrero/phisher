package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/adelapazborrero/gophisher/config"
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

	server := config.InitHTTP()
	server.Run()

}
