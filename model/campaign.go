package model

import (
	"database/sql"

	"github.com/adelapazborrero/gophisher/types"
)

type Campaign struct {
	ID    types.ID
	Title string
}

func NewCampaign(title string) *Campaign {
	return &Campaign{
		ID:    "",
		Title: title,
	}
}

func (c *Campaign) Find(db *sql.DB, id types.ID) (*Campaign, error) {
	var campaign Campaign
	row := db.QueryRow("SELECT * FROM campaigns WHERE id = $1", id)
	err := row.Scan(&campaign.ID, &campaign.Title)
	if err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (c *Campaign) FindAll(db *sql.DB) (*[]Campaign, error) {
	rows, err := db.Query("SELECT * FROM campaigns")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaignList []Campaign
	for rows.Next() {
		var campaign Campaign

		err = rows.Scan(&campaign.ID, &campaign.Title)
		if err != nil {
			return nil, err
		}
		campaignList = append(campaignList, campaign)
	}

	return &campaignList, nil
}

func (c *Campaign) Save(db *sql.DB) error {

	c.ID = types.NewID()

	_, err := db.Exec("INSERT INTO campaigns (id, title) VALUES ($1, $2)", c.ID, c.Title)
	if err != nil {
		return err
	}

	return nil
}
