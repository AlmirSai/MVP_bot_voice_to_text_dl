// bot/config/config.go
package config

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"tg-whisper-bot/bot/utils/logger"
)

// GetToken reads a specified .env file to retrieve the TELEGRAM_TOKEN value.
//
// It initializes a logger and logs the process of opening and reading the file.
// The function scans the file line by line, ignoring empty lines and comments,
// and splits each line into key-value pairs. If the key TELEGRAM_TOKEN is found,
// its corresponding value is returned. If the token is not found or if any error
// occurs during file operations or scanning, an error is returned.
//
// Parameters:
//   - filepath: Path to the .env file containing the TELEGRAM_TOKEN.
//
// Returns:
//   - The TELEGRAM_TOKEN value if found, otherwise an error.

func GetToken(filepath string) (string, error) {
	logFilePath := "storage/logs/config.log"

	if _, err := os.Stat("storage/logs"); os.IsNotExist(err) {
		err := os.MkdirAll("storage/logs", 0755)
		if err != nil {
			return "", fmt.Errorf("failed to create log directory: %v", err)
		}
	}

	logInstance, err := logger.GetLogger(logFilePath)
	if err != nil {
		return "", fmt.Errorf("Error initializing logger: %v", err)
	}
	defer logInstance.Close()

	logInstance.Info(fmt.Sprintf("Attempting to open file: %s", filepath))

	file, err := os.Open(filepath)
	if err != nil {
		logInstance.Error(fmt.Sprintf("Failed to open file: %v", err))
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			logInstance.Error(fmt.Sprintf("Failed to close file: %v", err))
		} else {
			logInstance.Info("File closed successfully")
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		logInstance.Info(fmt.Sprintf("Processing line: %s", line))

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			logInstance.Error(fmt.Sprintf("Invalid line in .env: %s", line))
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)

		if key == "TELEGRAM_TOKEN" {
			logInstance.Info("TELEGRAM_TOKEN found in .env")
			return value, nil
		}
	}

	if err := scanner.Err(); err != nil {
		logInstance.Error(fmt.Sprintf("Error reading .env: %v", err))
		return "", fmt.Errorf("error reading .env: %w", err)
	}

	logInstance.Error("TELEGRAM_TOKEN not found in .env")
	return "", fmt.Errorf("TELEGRAM_TOKEN not found in .env")
}

// CheckDependencies checks for the presence of required commands (ffmpeg, whisper)
// and log directories. If any required command or log directory is missing,
// an error is returned.
func CheckDependencies() error {
	logFilePath := "storage/logs/config.log"

	if _, err := os.Stat("storage/logs"); os.IsNotExist(err) {
		err := os.MkdirAll("storage/logs", 0755)
		if err != nil {
			return fmt.Errorf("failed to create log directory: %v", err)
		}
	}

	logInstance, err := logger.GetLogger(logFilePath)
	if err != nil {
		return fmt.Errorf("Error initializing logger: %v", err)
	}
	defer logInstance.Close()

	requiredCommands := []string{"ffmpeg", "whisper"}

	for _, cmd := range requiredCommands {
		logInstance.Info(fmt.Sprintf("Checking command: %s", cmd))
		if _, err := exec.LookPath(cmd); err != nil {
			logInstance.Error(fmt.Sprintf("Required command not found: %s", cmd))
			return fmt.Errorf("required command not found: %s", cmd)
		}
		logInstance.Info(fmt.Sprintf("Command %s found", cmd))
	}
	logInstance.Info("All required commands are present")
	return nil
}
