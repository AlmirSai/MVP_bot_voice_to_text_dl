package executils

import (
	"testing"
	"os/exec"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExecutor struct {
	mock.Mock
}

func (m *MockExecutor) LookPath(cmd string) (string, error) {
	args := m.Called(cmd)
	return args.String(0), args.Error(1)
}

func TestLookPath_CommandFound(t *testing.T) {
	mockExecutor := new(MockExecutor)
	mockExecutor.On("LookPath", "ffmpeg").Return("/usr/bin/ffmpeg", nil)

	ExecutorInstance = mockExecutor
	path, err := ExecutorInstance.LookPath("ffmpeg")

	assert.NoError(t, err)
	assert.Equal(t, "/usr/bin/ffmpeg", path)

	mockExecutor.AssertExpectations(t)
}

func TestLookPath_CommandNotFound(t *testing.T) {
	mockExecutor := new(MockExecutor)
	mockExecutor.On("LookPath", "nonexistentcmd").Return("", exec.ErrNotFound)


	ExecutorInstance = mockExecutor
	path, err := ExecutorInstance.LookPath("nonexistentcmd")

	assert.Error(t, err)
	assert.Equal(t, exec.ErrNotFound, err)
	assert.Empty(t, path)

	mockExecutor.AssertExpectations(t)
}

func TestLookPath_RealExecutor(t *testing.T) {
	ExecutorInstance = &RealExecutor{}
	path, err := ExecutorInstance.LookPath("ls")

	assert.NoError(t, err)
	assert.NotEmpty(t, path)
}
