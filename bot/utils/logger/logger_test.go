package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	logFilePathDebug   = "../../storage/logs/test_debug.log"
	logFilePathInfo    = "../../storage/logs/test_info.log"
	logFilePathWarning = "../../storage/logs/test_warning.log"
	logFilePathError   = "../../storage/logs/test_error.log"
	logFilePathFatal   = "../../storage/logs/test_fatal.log"
)

func ensureLogDirectoryExists() {
	if _, err := os.Stat("../../storage/logs"); os.IsNotExist(err) {
		err := os.MkdirAll("../../storage/logs", 0755)
		if err != nil {
			panic("Failed to create log directory: " + err.Error())
		}
	}
}

func TestLogger_Debug(t *testing.T) {
	ensureLogDirectoryExists() // Ensure the logs directory exists

	logInstance, err := GetLogger(logFilePathDebug)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()

	// Simulate logging a debug message
	logInstance.Debug("This is a debug message")

	// Check if log file is created
	_, err = os.Stat(logFilePathDebug)
	assert.NoError(t, err, "Expected log file to be created")
}

func TestLogger_Info(t *testing.T) {
	ensureLogDirectoryExists() // Ensure the logs directory exists

	logInstance, err := GetLogger(logFilePathInfo)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()

	// Simulate logging an info message
	logInstance.Info("This is an info message")

	// Check if log file is created
	_, err = os.Stat(logFilePathInfo)
	assert.NoError(t, err, "Expected log file to be created")
}

func TestLogger_Warning(t *testing.T) {
	ensureLogDirectoryExists() // Ensure the logs directory exists

	logInstance, err := GetLogger(logFilePathWarning)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()

	// Simulate logging a warning message
	logInstance.Warning("This is a warning message")

	// Check if log file is created
	_, err = os.Stat(logFilePathWarning)
	assert.NoError(t, err, "Expected log file to be created")
}

func TestLogger_Error(t *testing.T) {
	ensureLogDirectoryExists() // Ensure the logs directory exists

	logInstance, err := GetLogger(logFilePathError)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()

	// Simulate logging an error message
	logInstance.Error("This is an error message")

	// Check if log file is created
	_, err = os.Stat(logFilePathError)
	assert.NoError(t, err, "Expected log file to be created")
}

func TestLogger_Fatal(t *testing.T) {
	ensureLogDirectoryExists() // Ensure the logs directory exists

	logInstance, err := GetLogger(logFilePathFatal)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()

	// Simulate logging a fatal message
	logInstance.Fatal("This is a fatal message")

	// Check if log file is created
	_, err = os.Stat(logFilePathFatal)
	assert.NoError(t, err, "Expected log file to be created")
}
