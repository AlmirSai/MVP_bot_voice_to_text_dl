// bot/utils/executils/executils.go
package executils

import "os/exec"

type Executor interface {
	LookPath(cmd string) (string, error)
}

type RealExecutor struct{}

func (r *RealExecutor) LookPath(cmd string) (string, error) {
	return exec.LookPath(cmd)
}

var ExecutorInstance Executor = &RealExecutor{}
