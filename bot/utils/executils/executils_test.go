package executils

import (
	"testing"
	"os/exec"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockExecutor is a mock implementation of the Executor interface for testing.
type MockExecutor struct {
	mock.Mock
}

// LookPath mocks the LookPath function of the Executor interface.
func (m *MockExecutor) LookPath(cmd string) (string, error) {
	args := m.Called(cmd)
	return args.String(0), args.Error(1)
}

func TestLookPath_CommandFound(t *testing.T) {
	// Create a mock Executor
	mockExecutor := new(MockExecutor)

	// Define the expected behavior of the mock
	mockExecutor.On("LookPath", "ffmpeg").Return("/usr/bin/ffmpeg", nil)

	// Set the mockExecutor as the ExecutorInstance for testing
	ExecutorInstance = mockExecutor

	// Call the LookPath method with the command
	path, err := ExecutorInstance.LookPath("ffmpeg")

	// Assert the results
	assert.NoError(t, err)
	assert.Equal(t, "/usr/bin/ffmpeg", path)

	// Assert that the mock method was called with the expected argument
	mockExecutor.AssertExpectations(t)
}

func TestLookPath_CommandNotFound(t *testing.T) {
	// Create a mock Executor
	mockExecutor := new(MockExecutor)

	// Define the expected behavior of the mock
	mockExecutor.On("LookPath", "nonexistentcmd").Return("", exec.ErrNotFound)

	// Set the mockExecutor as the ExecutorInstance for testing
	ExecutorInstance = mockExecutor

	// Call the LookPath method with a command that doesn't exist
	path, err := ExecutorInstance.LookPath("nonexistentcmd")

	// Assert the results
	assert.Error(t, err)
	assert.Equal(t, exec.ErrNotFound, err)
	assert.Empty(t, path)

	// Assert that the mock method was called with the expected argument
	mockExecutor.AssertExpectations(t)
}

func TestLookPath_RealExecutor(t *testing.T) {
	// Set the real executor for testing
	ExecutorInstance = &RealExecutor{}

	// Call LookPath for a known command (e.g., "ls" or "ffmpeg")
	path, err := ExecutorInstance.LookPath("ls")

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that a path was returned
	assert.NotEmpty(t, path)
}
