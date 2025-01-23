package model

import "net/mail"

type MailMessage struct {
	From    mail.Address
	To      mail.Address
	Subject string
	Content string
}
