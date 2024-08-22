package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/adelapazborrero/gophisher/config"
	"github.com/adelapazborrero/gophisher/migration"
	"github.com/adelapazborrero/gophisher/routes"
	"github.com/adelapazborrero/gophisher/service/server"
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

	// reader := mail.NewCSVReader()
	// emails, err := reader.GetEmailsFromFile("./mail.txt")
	// if err != nil {
	// 	log.Fatal("Could not read emails from file. ", err)
	// }

	campaignService := server.NewCampaignService(db)

	server := config.NewServer(":8080", "localhost")

	config.RegisterRoutes(
		server.Engine,
		routes.CampaignRoutes(campaignService),
	)

	log.Printf("Routes registered")

	server.Start()

}
