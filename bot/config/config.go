package config

import (
	"fmt"
	"os/exec"
)

func GetToken() string {
	return "TELEGRAM_TOKEN"
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