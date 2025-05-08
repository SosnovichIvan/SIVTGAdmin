package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Интерфейс логгера
type Logger interface {
	Info(prefix string, msg string)
	Warn(prefix string, msg string)
	Error(prefix string, msg string)
	Fatal(prefix string, msg string)
}

// Реализация логгера поверх logrus
type AppLogger struct {
	*logrus.Logger
}

func NewLogger(debug_mode bool) Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	file, _ := os.Create("output.log")
	log.Out = io.MultiWriter(file, os.Stdout)

	if debug_mode {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	return &AppLogger{
		Logger: log,
	}
}

func (l *AppLogger) Info(prefix string, msg string) {
	l.WithField("service: ", prefix).Info(msg)
}

func (l *AppLogger) Warn(prefix string, msg string) {
	l.WithField("service: ", prefix).Warn(msg)
}

func (l *AppLogger) Error(prefix string, msg string) {
	l.WithField("service: ", prefix).Error(msg)
}

func (l *AppLogger) Fatal(prefix string, msg string) {
	l.WithField("service: ", prefix).Fatal(msg)
}
