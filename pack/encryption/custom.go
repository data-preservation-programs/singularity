package encryption

import (
	"github.com/pkg/errors"
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
}

func (rc *readCloserWithError) Read(p []byte) (n int, err error) {
	n, err = rc.ReadCloser.Read(p)
	if errors.Is(err, io.EOF) {
		rc.wait.Wait()
		if *rc.err != nil {
			return 0, *rc.err
		}
	}
	return n, err
}

func (rc *readCloserWithError) Close() error {
	return rc.ReadCloser.Close()
}

func (c CustomEncryptor) Encrypt(in io.Reader) (io.ReadCloser, error) {
	// Invoke the underlying command, use in as stdin, and return the stdout as a ReadCloser
	// Set the input
	c.cmd.Stdin = in

	// Start the command and get the output pipe
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

	return &readCloserWithError{stdout, waitErr, wg}, nil
}

func NewCustomEncryptor(cmd *exec.Cmd) Encryptor {
	return &CustomEncryptor{cmd: cmd}
}
