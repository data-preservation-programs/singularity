package api

import "bytes"

// MockWorkerCmd implements WorkerCmd for testing
type MockWorkerCmd struct {
	running bool
	pid     int
	stdout  *bytes.Buffer
	stderr  *bytes.Buffer
}

func NewMockWorkerCmd() *MockWorkerCmd {
	return &MockWorkerCmd{
		pid: 12345, // Fixed test PID
	}
}

func (m *MockWorkerCmd) Start() error {
	m.running = true
	return nil
}

func (m *MockWorkerCmd) Kill() error {
	m.running = false
	return nil
}

func (m *MockWorkerCmd) Pid() int {
	if !m.running {
		return 0
	}
	return m.pid
}

func (m *MockWorkerCmd) SetOutput(stdout, stderr *bytes.Buffer) {
	m.stdout = stdout
	m.stderr = stderr
}
