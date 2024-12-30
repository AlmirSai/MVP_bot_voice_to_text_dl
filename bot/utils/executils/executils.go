// bot/utils/executils/executils.go
package executils

import "os/exec"

type Executor interface {
	LookPath(cmd string) (string, error)
}

type RealExecutor struct{}

// LookPath implements the Executor interface by calling os/exec.LookPath.
// It returns the path for the given command and an error if the command is not found.
func (r *RealExecutor) LookPath(cmd string) (string, error) {
	return exec.LookPath(cmd)
}

var ExecutorInstance Executor = &RealExecutor{}
