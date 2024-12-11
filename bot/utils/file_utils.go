package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"tg-whisper-bot/bot/utils/logger"
)

func DownloadFile(filePath, url string) error {
	logFilePath := "storage/logs/file_utils.log"

	if _, err := os.Stat("storage/logs"); os.IsNotExist(err) {
		err := os.MkdirAll("storage/logs", 0755)
		if err != nil {
			fmt.Println("Failed to create log directory:", err)
			return err
		}
	}

	logInstance, err := logger.GetLogger(logFilePath)
	if err != nil {
		fmt.Println("Error initializing logger:", err)
		return err
	}
	defer logInstance.Close()

	logInstance.Info(fmt.Sprintf("Started downloading file from URL: %s to path: %s", url, filePath))

	resp, err := http.Get(url)
	if err != nil {
		logInstance.Error(fmt.Sprintf("Failed to get URL %s: %v", url, err))
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		logInstance.Error(fmt.Sprintf("Failed to create file %s: %v", filePath, err))
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logInstance.Error(fmt.Sprintf("Failed to copy data from URL %s to file %s: %v", url, filePath, err))
		return err
	}

	logInstance.Info(fmt.Sprintf("Successfully downloaded file from URL: %s to path: %s", url, filePath))

	return nil
}
