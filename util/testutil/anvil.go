package testutil

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type AnvilInstance struct {
	RPCURL string
	cmd    *exec.Cmd
}

func (a *AnvilInstance) Close() {
	if a.cmd != nil && a.cmd.Process != nil {
		a.cmd.Process.Kill()
		a.cmd.Wait()
	}
}

// StartAnvil forks the given upstream RPC on a random free port.
// Returns when the fork is ready to accept connections.
func StartAnvil(t *testing.T, upstreamRPC string) *AnvilInstance {
	t.Helper()

	if _, err := exec.LookPath("anvil"); err != nil {
		t.Skip("anvil not found on PATH")
	}

	port, err := freePort()
	if err != nil {
		t.Fatalf("finding free port: %v", err)
	}

	cmd := exec.Command("anvil",
		"--fork-url", upstreamRPC,
		"--port", fmt.Sprintf("%d", port),
		"--silent",
	)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("starting anvil: %v", err)
	}

	rpcURL := fmt.Sprintf("http://127.0.0.1:%d", port)
	inst := &AnvilInstance{RPCURL: rpcURL, cmd: cmd}
	t.Cleanup(inst.Close)

	if err := waitForRPC(rpcURL, 30*time.Second); err != nil {
		t.Fatalf("anvil not ready: %v", err)
	}

	return inst
}

func freePort() (int, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	port := l.Addr().(*net.TCPAddr).Port
	l.Close()
	return port, nil
}

func waitForRPC(url string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		client, err := ethclient.DialContext(ctx, url)
		if err == nil {
			_, err = client.ChainID(ctx)
			client.Close()
			cancel()
			if err == nil {
				return nil
			}
		} else {
			cancel()
		}
		time.Sleep(200 * time.Millisecond)
	}
	return fmt.Errorf("rpc at %s not ready after %s", url, timeout)
}
