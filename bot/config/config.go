package config

import (
	"bufio"
	"fmt"	
	"os"
	"os/exec"
	"strings"
)

func GetToken(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid line in .env: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)

		if key == "TELEGRAM_TOKEN" {
			return value, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading .env: %w", err)
	}

	return "", fmt.Errorf("TELEGRAM_TOKEN not found in .env")
}

func CheckDependencies() error {
	requiredCommands := []string{"ffmpeg", "whisper"}
	
	for _, cmd := range requiredCommands {
		if _, err := exec.LookPath(cmd); err != nil {
			return fmt.Errorf("requiered command not found: %s", cmd)
		}
	}
	return nil
}