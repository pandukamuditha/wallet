package common

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() *Logger {
	return &Logger{logger: log.New(os.Stdout, "blog-api ", log.LstdFlags)}
}

func (l *Logger) Log(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *Logger) Logf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}
