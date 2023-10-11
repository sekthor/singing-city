package repo

import (
	"fmt"

	"github.com/sekthor/songbird-backend/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(conf config.DbConfig) (*gorm.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", 
        conf.User, conf.Pass, conf.Host, conf.Port, conf.Database, "charset=utf8mb4&parseTime=True&loc=Local")
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}


