package server

import (
	"database/sql"

	"github.com/adelapazborrero/gophisher/model"
)

type IMailService interface {
	Get() (*model.Mail, error)
	GetAll() (*[]model.Mail, error)
	Create() error
	Update() error
	Delete() error
}

type MailService struct {
	db *sql.DB
}

func NewMailService(db *sql.DB) IMailService {
	return &MailService{
		db: db,
	}
}

func (m *MailService) Get() (*model.Mail, error) {
	return nil, nil
}

func (m *MailService) GetAll() (*[]model.Mail, error) {
	return nil, nil
}

func (m *MailService) Create() error {
	return nil
}

func (m *MailService) Update() error {
	return nil
}

func (m *MailService) Delete() error {
	return nil
}
