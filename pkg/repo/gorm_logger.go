package repo

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
)

type GormLogger struct{}

func (l GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	log.Info().Msg(msg)
}

func (l GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	log.Warn().Msg(msg)
}

func (l GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	log.Error().Msg(msg)
}

func (GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	var msg string
	sql, rows := fc()

	if err != nil {
		msg = "sql query error"
	} else {
		msg = "sql query success"
	}

	log.Trace().Err(err).
		Str("query", sql).
		Int64("rows", rows).
		Msg(msg)
}
