package service

import (
	"bytes"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/rs/zerolog/log"
	"github.com/sekthor/singing-city/pkg/config"
	"github.com/sekthor/singing-city/pkg/model"
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

func (s *NotificationService) SendApplicationMessage(recipient string, params MessageParams) error {

	// do not send mails unless it is enabled in the config
	if !s.conf.EnableMail {
		return nil
	}

	var msg bytes.Buffer
	if err := applicationMessageTmpl.Execute(&msg, params); err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("could not template application email for %s", recipient))
		return err
	}
	err := s.Send(ApplicationMessageSubject, msg.String(), recipient)

	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("could not send application email to %s", recipient))
		return err
	}

	log.Info().Msg(fmt.Sprintf("sent register email to %s", recipient))
	return nil
}

func (s *NotificationService) SendConfirmedMessage(recipient string, params MessageParams) error {

	// do not send mails unless it is enabled in the config
	if !s.conf.EnableMail {
		return nil
	}

	var msg bytes.Buffer
	if err := confirmedMessageTmpl.Execute(&msg, params); err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("could not template confirmed email for %s", recipient))
		return err
	}
	err := s.Send(ConfirmedMessageSubject, msg.String(), recipient)

	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("could not send confirmed email to %s", recipient))
		return err
	}

	log.Info().Msg(fmt.Sprintf("sent confirmed email to %s", recipient))
	return nil
}

func (s *NotificationService) SendRejectedMessage(recipient string, params MessageParams) error {

	// do not send mails unless it is enabled in the config
	if !s.conf.EnableMail {
		return nil
	}

	var msg bytes.Buffer
	if err := rejectedMessageTmpl.Execute(&msg, params); err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("could not template rejected email for %s", recipient))
		return err
	}

	err := s.Send(RejectedMessageSubject, msg.String(), recipient)

	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("could not send rejected email to %s", recipient))
		return err
	}

	log.Info().Msg(fmt.Sprintf("sent rejected email to %s", recipient))
	return nil
}

func (s *NotificationService) SendPasswordResetLink(recipient string, params MessageParams) error {
	// do not send mails unless it is enabled in the config
	if !s.conf.EnableMail {
		return nil
	}

	var msg bytes.Buffer
	if err := passwordResetMessage.Execute(&msg, params); err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("could not template password reset email for %s", recipient))
		return err
	}

	err := s.Send(PasswordResetSubject, msg.String(), recipient)

	if err != nil {
		log.Warn().Err(err).Msg(fmt.Sprintf("could not send password reset email to %s", recipient))
		return err
	}

	log.Info().Msg(fmt.Sprintf("sent password reset email to %s", recipient))
	return nil
}
