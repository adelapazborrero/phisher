package model

import (
	"database/sql"
	"time"

	"github.com/adelapazborrero/gophisher/types"
)

type Mail struct {
	ID       types.ID
	Sender   string
	Receiver string
	SentTime *time.Time
	OpenTime *time.Time
	Clicked  bool
}

func NewMail(from string, to string) *Mail {
	return &Mail{
		ID:       "",
		Sender:   from,
		Receiver: to,
		SentTime: types.TimePtr(time.Now()),
		OpenTime: nil,
		Clicked:  false,
	}
}

func (m *Mail) Find(db *sql.DB, id types.ID) (*Mail, error) {
	var mail Mail
	row := db.QueryRow("SELECT * FROM mails WHERE id = $1", id)
	err := row.Scan(&mail.ID, &mail.Sender, &mail.Receiver, &mail.SentTime, &mail.OpenTime, &mail.Clicked)
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

		err := rows.Scan(&mail.ID, &mail.Sender, &mail.Receiver, &mail.SentTime, &mail.OpenTime, &mail.Clicked)
		if err != nil {
			return nil, err
		}
		mailList = append(mailList, mail)
	}

	return &mailList, nil
}

func (m *Mail) Save(db *sql.DB) error {

	m.ID = types.NewID()

	_, err := db.Exec(`
		INSERT INTO mails (id, sender, receiver, sent_time, open_time) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`,
		m.ID,
		m.Sender,
		m.Receiver,
		m.SentTime,
		m.OpenTime,
		m.Clicked,
	)
	if err != nil {
		return err
	}

	return nil
}
