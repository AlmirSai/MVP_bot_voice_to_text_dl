// bot/utils/logger/logger.go
package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger represents a logger with a file and a log.Logger instance.
type Logger struct {
	file   *os.File
	logger *log.Logger
}

// NewLogger creates a new Logger instance with the specified file path.
func NewLogger(filepath string) (*Logger, error) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile)

	return &Logger{
		file:   file,
		logger: logger,
	}, nil
}

// GetLogger initializes the logging directory and returns a new Logger instance.
func GetLogger(logFilePath string) (*Logger, error) {
	if _, err := os.Stat("storage/logs"); os.IsNotExist(err) {
		if err := os.MkdirAll("storage/logs", 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %w", err)
		}
	}

	return NewLogger(logFilePath)
}

// Debug logs a debug message.
func (l *Logger) Debug(message string) {
	l.logger.SetPrefix("DEBUG: ")
	l.logger.Println(message)
}

// Trace logs a trace message.
func (l *Logger) Trace(message string) {
	l.logger.SetPrefix("TRACE: ")
	l.logger.Println(message)
}

// Info logs an informational message.
func (l *Logger) Info(message string) {
	l.logger.SetPrefix("INFO: ")
	l.logger.Println(message)
}

// Warning logs a warning message.
func (l *Logger) Warning(message string) {
	l.logger.SetPrefix("WARNING: ")
	l.logger.Println(message)
}

// Error logs an error message.
func (l *Logger) Error(message string) {
	l.logger.SetPrefix("ERROR: ")
	l.logger.Println(message)
}

// Fatal logs a fatal error message.
func (l *Logger) Fatal(message string) {
	l.logger.SetPrefix("FATAL: ")
	l.logger.Println(message)
}

// Close closes the log file.
func (l *Logger) Close() error {
	return l.file.Close()
}
