package server

import (
	"database/sql"

	"github.com/adelapazborrero/gophisher/model"
	"github.com/adelapazborrero/gophisher/types"
)

type CampaignService struct {
	db *sql.DB
}

func NewCampaignService(db *sql.DB) CampaignService {
	return CampaignService{
		db: db,
	}
}

func (m *CampaignService) Get(id types.ID) (*model.Campaign, error) {
	campaign := &model.Campaign{}
	return campaign.Find(m.db, id)
}

func (m *CampaignService) GetAll() (*[]model.Campaign, error) {
	campaign := &model.Campaign{}
	return campaign.FindAll(m.db)
}

func (m *CampaignService) Create(title string) error {
	campaign := model.NewCampaign(title)
	return campaign.Save(m.db)
}

func (m *CampaignService) Update(id types.ID, title string) error {
	campaign := &model.Campaign{}
	return campaign.Update(m.db, id, title)
}

func (m *CampaignService) Delete(id types.ID) error {
	campaign := &model.Campaign{}
	return campaign.Delete(m.db, id)
}
