package repo

import (
	"errors"
	"fmt"
	"os"

	"github.com/sekthor/songbird-backend/pkg/config"
	"github.com/sekthor/songbird-backend/pkg/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(conf config.DbConfig) (*gorm.DB, error) {

	switch conf.Type {
	case "sqlite":
		// make sure the sqlite database file exits
		if _, err := os.Stat(conf.Database); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				file, err := os.Create(conf.Database)
				if err != nil {
					return nil, err
				}
				file.Close()
			}
			return nil, err
		}
		return gorm.Open(sqlite.Open(conf.Database), &gorm.Config{
			TranslateError: false,
			Logger:         GormLogger{},
		})
	default:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			conf.User, conf.Pass, conf.Host, conf.Port, conf.Database, "charset=utf8mb4&parseTime=True&loc=Local")
		return gorm.Open(mysql.Open(dsn), &gorm.Config{
			TranslateError: false,
			Logger:         GormLogger{},
		})
	}
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Venue{},
		&model.Artist{},
		&model.Timeslot{},
		&model.Application{},
	)
}
