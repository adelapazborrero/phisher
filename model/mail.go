package model

import (
	"database/sql"
	"time"

	"github.com/adelapazborrero/gophisher/types"
)

type Mail struct {
	ID       types.ID
	From     string
	To       string
	SentTime *time.Time
	OpenTime *time.Time
}

func NewMail(from string, to string) *Mail {
	return &Mail{
		ID:       "",
		From:     from,
		To:       to,
		SentTime: types.TimePtr(time.Now()),
		OpenTime: nil,
	}
}

func (m *Mail) Find(db *sql.DB, id types.ID) (*Mail, error) {
	var mail Mail
	row := db.QueryRow("SELECT * FROM mails WHERE id = $1", id)
	err := row.Scan(&mail.ID, &mail.From, &mail.To, &mail.SentTime, &mail.OpenTime)
	if err != nil {
		return nil, err
	}
	return &mail, nil
}

func (m *Mail) FindAll(db *sql.DB) (*[]Mail, error) {
	rows, err := db.Query("SELECT * FROM mails")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mailList []Mail
	for rows.Next() {
		var mail Mail

		err := rows.Scan(&mail.ID, &mail.From, &mail.To, &mail.SentTime, &mail.OpenTime)
		if err != nil {
			return nil, err
		}
		mailList = append(mailList, mail)
	}

	return &mailList, nil
}

func (m *Mail) Save(db *sql.DB) error {

	m.ID = types.NewID()

	_, err := db.Exec("INSERT INTO mails (id, from, to, sent_time, open_time) VALUES ($1, $2, $3, $4, $5)", m.ID, m.From, m.To, m.SentTime, m.OpenTime)
	if err != nil {
		return err
	}

	return nil
}
