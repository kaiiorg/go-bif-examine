package gorm_logger

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	log      zerolog.Logger
	LogLevel logger.LogLevel
}

func NewGormLogger(log zerolog.Logger, level zerolog.Level) GormLogger {
	gl := GormLogger{
		log: log,
	}

	switch level {
	case zerolog.TraceLevel:
		gl.LogLevel = logger.Info
	case zerolog.DebugLevel:
		gl.LogLevel = logger.Info
	case zerolog.InfoLevel:
		gl.LogLevel = logger.Info
	case zerolog.WarnLevel:
		gl.LogLevel = logger.Warn
	case zerolog.ErrorLevel:
		gl.LogLevel = logger.Error
	case zerolog.FatalLevel:
		gl.LogLevel = logger.Error
	case zerolog.PanicLevel:
		gl.LogLevel = logger.Error
	default:
		gl.LogLevel = logger.Info
	}

	return gl
}

func (gl GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return &gl
}

func (gl GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	gl.genericLog(zerolog.InfoLevel, msg, data)
}

func (gl GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	gl.genericLog(zerolog.WarnLevel, msg, data)
}

func (gl GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	gl.genericLog(zerolog.ErrorLevel, msg, data)
}

func (gl GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	log.Debug().
		Err(err).
		Dur("elapsed", elapsed).
		Int64("rows", rows).
		Str("sql", sql).
		Msg("query")
}

func (gl GormLogger) genericLog(level zerolog.Level, msg string, data ...interface{}) {
	event := gl.log.WithLevel(level)
	for i, d := range data {
		event = event.Interface(
			fmt.Sprintf("data%d", i),
			d,
		)
	}
	event.Msg(msg)
}
