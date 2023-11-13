package service

import (
	"bytes"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/rs/zerolog/log"
	"github.com/sekthor/songbird-backend/pkg/config"
	"github.com/sekthor/songbird-backend/pkg/model"
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

func (s *NotificationService) SendRegisterMessage(user model.User) error {

	// do not send mails unless it is enabled in the config
	if !s.conf.EnableMail {
		return nil
	}

	var msg bytes.Buffer
	if err := registerMessageTmpl.Execute(&msg, user); err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("could not template register email for %s", user.Email))
		return err
	}
	err := s.Send(RegisterMessageSubject, msg.String(), user.Email)

	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("could not send register email to %s", user.Email))
		return err
	}

	log.Info().Msg(fmt.Sprintf("sent register email to %s", user.Email))
	return nil
}
