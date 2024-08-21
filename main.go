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

	_, err = db.Exec(migration.CreateMailTable)
	if err != nil {
		log.Fatal("Could not create table:", err)
	}

	campaign := model.NewCampaign("My new campaign")
	campaigns, err := campaign.Find(db, "b160c3f8-b18b-4695-8631-f53eefd530b9")
	if err != nil {
		log.Fatal("QUERY ERROR -> ", err)
	}
	fmt.Println(campaigns)

	mail := model.NewMail("abel45991690@gmail.com", "test@mail.com")
	err = mail.Save(db)
	if err != nil {
		log.Fatal("Could not create mail.", err)
	}

	mails, err := mail.FindAll(db)
	if err != nil {
		log.Fatal("Could not get emails", err)
	}
	fmt.Println(mails)

}
