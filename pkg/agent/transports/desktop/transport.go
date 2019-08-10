package desktop

import (
	"os"
	"os/exec"
)

type Transport struct{}

func (Transport) Copy(localPath, remoteName string) error {
	panic("implement me")
}

func (Transport) Command(command string) (*exec.Cmd, error) {
	panic("implement me")
}

func (Transport) ClassifyError(processState *os.ProcessState, errorOutput string) (bool, bool, error) {
	panic("implement me")
}
