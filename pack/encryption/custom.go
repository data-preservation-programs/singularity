package encryption

import (
	"io"
	"os/exec"
	"sync"
)

// CustomEncryptor is an implementation of Encryptor that uses a custom command to encrypt data.
type CustomEncryptor struct {
	cmd *exec.Cmd
}

type readCloserWithError struct {
	io.ReadCloser
	err  *error
	wait *sync.WaitGroup
	cmd  *exec.Cmd
}

func (rc *readCloserWithError) Read(p []byte) (n int, err error) {
	return rc.ReadCloser.Read(p)
}

func (rc *readCloserWithError) Close() error {
	err := rc.ReadCloser.Close()
	rc.wait.Wait()
	if *rc.err != nil {
		return *rc.err
	}
	return err
}

func (c CustomEncryptor) Encrypt(in io.Reader) (io.ReadCloser, error) {
	c.cmd.Stdin = in

	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	// Start the command
	err = c.cmd.Start()
	if err != nil {
		return nil, err
	}

	waitErr := new(error)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		*waitErr = c.cmd.Wait()
	}()

	return &readCloserWithError{stdout, waitErr, wg, c.cmd}, nil
}

func NewCustomEncryptor(cmd *exec.Cmd) Encryptor {
	return &CustomEncryptor{cmd: cmd}
}
