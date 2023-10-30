package service

import (
	"fmt"
	"net/smtp"

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

func (s *NotificationService) send(msg string, recipients ...string) error {
	auth := smtp.PlainAuth("", s.conf.Email, s.conf.Password, s.conf.Server)
	return smtp.SendMail(
		fmt.Sprintf("%s:%s", s.conf.Server, s.conf.Port),
		auth,
		"Singing City",
		recipients,
		[]byte(msg),
	)
}
