package config

import (
	"os"

	"github.com/rs/zerolog"
)

type Config struct {
	Server ServerConfig
	DB     DbConfig
	Smtp   SmtpConfig
}

type ServerConfig struct {
	Host     string
	Port     string
	Secret   string
	Loglevel string
}

type DbConfig struct {
	Type     string
	Database string
	Host     string
	Port     string
	User     string
	Pass     string
}

type SmtpConfig struct {
	SenderEmail string
	SenderName  string
	Password    string
	Server      string
	Port        string
	EnableMail  string
}

func LoadConfig() Config {
	conf := Config{}
	conf.Server.Host = os.Getenv("SERVERHOST")
	conf.Server.Port = os.Getenv("SERVERPORT")
	conf.Server.Secret = os.Getenv("SERVERSECRET")
	conf.Server.Loglevel = os.Getenv("SERVERLOGLEVEL")
	conf.DB.Type = os.Getenv("DBTYPE")
	conf.DB.Database = os.Getenv("DBDATABASE")
	conf.DB.Host = os.Getenv("DBHOST")
	conf.DB.Port = os.Getenv("DBPORT")
	conf.DB.User = os.Getenv("DBUSER")
	conf.DB.Pass = os.Getenv("DBPASS")
	conf.Smtp.Password = os.Getenv("SMTPPASSWORD")
	conf.Smtp.SenderName = os.Getenv("SMTPSENDERNAME")
	conf.Smtp.SenderEmail = os.Getenv("SMTPSENDEREMAIL")
	conf.Smtp.Port = os.Getenv("SMTPSERVERPORT")
	conf.Smtp.Server = os.Getenv("SMTPSERVER")
	return conf
}

func (c ServerConfig) GetLoglevel() zerolog.Level {
	switch c.Loglevel {
	case "panic":
		return zerolog.PanicLevel
	case "fatal":
		return zerolog.FatalLevel
	case "error":
		return zerolog.ErrorLevel
	case "warn":
		return zerolog.WarnLevel
	case "info":
		return zerolog.InfoLevel
	case "debug":
		return zerolog.DebugLevel
	case "trace":
		return zerolog.TraceLevel
	default:
		return zerolog.InfoLevel
	}
}
