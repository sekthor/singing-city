package config

import "os"


type Config struct {
    Server ServerConfig
    DB DbConfig
}

type ServerConfig struct {
    Host string
    Port string
}

type DbConfig struct {
    Database string
    Host string
    Port string
    User string
    Pass string
}


func LoadConfig() Config {
    conf := Config{}
    conf.Server.Host = os.Getenv("SERVERHOST")
    conf.Server.Port = os.Getenv("SERVERPORT")
    conf.DB.Database = os.Getenv("DBDATABASE")
    conf.DB.Host = os.Getenv("DBHOST")
    conf.DB.Port = os.Getenv("DBPORT")
    conf.DB.User = os.Getenv("DBUSER")
    conf.DB.Pass= os.Getenv("DBPASS")
    return conf
}
