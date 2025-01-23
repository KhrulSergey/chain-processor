package service

import (
	"strings"

	"github.com/khrulsergey/chain-processor/pkg/model"
	"go.uber.org/zap"
)

const (
	CRLF = "\r\n"
)

type MailService interface {
	SendMail(model.MailMessage) error
}

var (
	_ MailService = &mockMailService{}
)

type mockMailService struct {
	log *zap.SugaredLogger
}

func NewMockMailService(log *zap.SugaredLogger) *mockMailService {
	return &mockMailService{
		log: log,
	}
}

func (s *mockMailService) SendMail(message model.MailMessage) error {
	rawMail := createRawMessage(message)
	s.log.Infof("mock mail service sending a message:\n%s", rawMail)
	return nil
}

func createRawMessage(message model.MailMessage) string {
	b := strings.Builder{}

	//Headers
	b.WriteString("From: ")
	b.WriteString(message.From.String())
	b.WriteString(CRLF)

	b.WriteString("To: ")
	b.WriteString(message.To.String())
	b.WriteString(CRLF)

	b.WriteString("Subject: ")
	b.WriteString(message.Subject)
	b.WriteString(CRLF)

	b.WriteString("Content-Type: text/plain; charset=\"utf-8\"")
	b.WriteString(CRLF)

	b.WriteString(CRLF)

	//Body
	b.WriteString(message.Content)

	return b.String()
}
