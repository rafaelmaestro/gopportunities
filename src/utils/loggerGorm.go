package utils

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
)

type LoggerGorm struct {
	Log zerolog.Logger
}

func (zl LoggerGorm) LogMode(level logger.LogLevel) logger.Interface {
	return zl
}

func (zl LoggerGorm) Info(ctx context.Context, msg string, data ...interface{}) {
	zl.Log.Info().Msgf(msg, data...)
}

func (zl LoggerGorm) Warn(ctx context.Context, msg string, data ...interface{}) {
	zl.Log.Warn().Msgf(msg, data...)
}

func (zl LoggerGorm) Error(ctx context.Context, msg string, data ...interface{}) {
	zl.Log.Error().Msgf(msg, data...)
}

func (zl LoggerGorm) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	duration := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		zl.Log.Error().
			Err(err).
			Dur("duration", duration).
			Str("sql", sql).
			Int64("rows", rows).
			Msg("Error executing query")
	} else {
		zl.Log.Info().
			Dur("duration", duration).
			Str("sql", sql).
			Int64("rows", rows).
			Msg("Query executed")
	}
}
