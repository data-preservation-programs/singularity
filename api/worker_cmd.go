package api

import (
	"bytes"
	"os/exec"
)

// WorkerCmd abstracts process control operations
type WorkerCmd interface {
	Start() error
	Kill() error
	Pid() int
	SetOutput(stdout, stderr *bytes.Buffer)
}

// RealWorkerCmd wraps exec.Cmd for real process execution
type RealWorkerCmd struct {
	cmd *exec.Cmd
}

func NewRealWorkerCmd(name string, args ...string) *RealWorkerCmd {
	return &RealWorkerCmd{
		cmd: exec.Command(name, args...),
	}
}

func (r *RealWorkerCmd) Start() error {
	return r.cmd.Start()
}

func (r *RealWorkerCmd) Kill() error {
	if r.cmd.Process == nil {
		return nil
	}
	return r.cmd.Process.Kill()
}

func (r *RealWorkerCmd) Pid() int {
	if r.cmd.Process == nil {
		return 0
	}
	return r.cmd.Process.Pid
}

func (r *RealWorkerCmd) SetOutput(stdout, stderr *bytes.Buffer) {
	r.cmd.Stdout = stdout
	r.cmd.Stderr = stderr
}
