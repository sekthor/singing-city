package service

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/sekthor/songbird-backend/pkg/config"
)

type NotificationService struct {
	conf config.SmtpConfig
}

func NewNotificationService(conf config.SmtpConfig) NotificationService {
	return NotificationService{
		conf: conf,
	}
}

func (s *NotificationService) Send(subject string, msg string, recipients ...string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", s.conf.SenderName, s.conf.SenderEmail)
	e.Subject = subject
	e.HTML = []byte(msg)
	e.To = recipients

	auth := smtp.PlainAuth("", s.conf.SenderEmail, s.conf.Password, s.conf.Server)
	host := fmt.Sprintf("%s:%s", s.conf.Server, s.conf.Port)

	return e.Send(host, auth)
}
