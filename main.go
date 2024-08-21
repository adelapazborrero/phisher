package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/adelapazborrero/gophisher/migration"
	"github.com/adelapazborrero/gophisher/model"
)

func main() {
	db, err := sql.Open("sqlite3", "gophisher.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(migration.CreateCampaignTable)
	if err != nil {
		log.Fatal("Could not create table:", err)
	}

	campaign := model.NewCampaign("My new campaign")
	campaigns, err := campaign.Find(db, "b160c3f8-b18b-4695-8631-f53eefd530b9")
	if err != nil {
		log.Fatal("QUERY ERROR -> ", err)
	}

	fmt.Println(campaigns)
}
