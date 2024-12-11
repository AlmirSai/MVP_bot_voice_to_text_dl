package logger

import (
	"log"
	"os"
)

type Logger struct {
	file   *os.File
	logger *log.Logger
}

func NewLogger(filepath string) (*Logger, error) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)

	return &Logger{
		file:   file,
		logger: logger,
	}, nil
}

func (l *Logger) Error(message string) {
	l.logger.SetPrefix("ERROR: ")
	l.logger.Println(message)
}

func (l *Logger) Fatal(message string) {
	l.logger.SetPrefix("FATAL: ")
	l.logger.Println(message)
}

func (l *Logger) Info(message string) {
	l.logger.SetPrefix("INFO: ")
	l.logger.Println(message)
}

func (l *Logger) Close() error {
	return l.file.Close()
}
