package mail

import "net/smtp"

type IMailer interface {
	Send() error
}

type Mailer struct {
	From     string
	To       []string
	Password string
	smtpHost string
	smtpPort string
}

func NewMailer(from string, to []string, password string) IMailer {
	return &Mailer{
		From:     from,
		To:       to,
		Password: password,
		smtpHost: "smtp.gmail.com",
		smtpPort: "587",
	}
}

func (m *Mailer) Send() error {
	message := []byte("My super secret message.")
	auth := smtp.PlainAuth("", m.From, m.Password, m.smtpHost)

	err := smtp.SendMail(m.smtpHost+":"+m.smtpPort, auth, m.From, m.To, message)
	if err != nil {
		return err
	}

	return nil
}
