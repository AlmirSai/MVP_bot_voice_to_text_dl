// bot/config/config_test.go
package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"tg-whisper-bot/bot/utils/logger"
	"tg-whisper-bot/bot/utils/executils"
)

const (
	logFilePathGetTokenSuccess      = "../../storage/logs/test_get_token.log"
	logFilePathGetTokenFileNotFound = "../../storage/logs/test_get_token_file_not_found.log"
	logFilePathGetTokenTokenNotFound = "../../storage/logs/test_get_token_token_not_found.log"
	logFilePathCheckDepsSuccess     = "../../storage/logs/test_check_dependencies_success.log"
	logFilePathCheckDepsFail        = "../../storage/logs/test_check_dependencies_fail.log"
)

type MockExecutor struct {
	mock.Mock
}

func (m *MockExecutor) LookPath(cmd string) (string, error) {
	args := m.Called(cmd)
	return args.String(0), args.Error(1)
}

func TestGetToken_Success(t *testing.T) {
	logInstance, err := logger.GetLogger(logFilePathGetTokenSuccess)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()

	testFile := "test.env"
	f, err := os.Create(testFile)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFile)
	defer f.Close()

	f.WriteString("TELEGRAM_TOKEN=1234567890")

	token, err := GetToken(testFile)
	assert.NoError(t, err)
	assert.Equal(t, "1234567890", token)

	logInstance.Info("Successfully retrieved token")
}

func TestGetToken_FileNotFound(t *testing.T) {
	logInstance, err := logger.GetLogger(logFilePathGetTokenFileNotFound)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()

	token, err := GetToken("nonexistent.env")
	assert.Error(t, err)
	assert.Empty(t, token)

	logInstance.Error("File not found")
}

func TestGetToken_TokenNotFound(t *testing.T) {
	logInstance, err := logger.GetLogger(logFilePathGetTokenTokenNotFound)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()

	testFile := "test_no_token.env"
	f, err := os.Create(testFile)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFile)
	defer f.Close()

	f.WriteString("SOME_OTHER_KEY=some_value")

	token, err := GetToken(testFile)
	assert.Error(t, err)
	assert.Empty(t, token)

	logInstance.Error("Token not found")
}

func TestCheckDependencies_Success(t *testing.T) {
	mockExecutor := new(MockExecutor)
	executils.ExecutorInstance = mockExecutor

	mockExecutor.On("LookPath", "ffmpeg").Return("/usr/bin/ffmpeg", nil)
	mockExecutor.On("LookPath", "whisper").Return("/usr/bin/whisper", nil)

	logInstance, err := logger.GetLogger(logFilePathCheckDepsSuccess)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()

	err = CheckDependencies()
	assert.NoError(t, err)

	logInstance.Info("Dependencies check successful")
}
