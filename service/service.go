package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"

	"github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)

type Done = <-chan struct{}

type Fail = <-chan error

// Server is an interface that represents a server that can be started and has a name.
type Server interface {
	// Start Starts the server and returns two channels. The Done channel is closed when the server is done running,
	//	and the Fail channel receives an error if the server fails. The method also returns an error immediately
	//	if the server fails to start. The server runs until the provided context is cancelled.
	Start(ctx context.Context) (Done, Fail, error)
	// Name returns the name of the server.
	Name() string
}

var ErrNoService = errors.New("no service is running")

// StartServers is a function that starts multiple servers concurrently.
// It takes a context, a logger, and a variadic slice of servers as arguments.
// If no servers are provided, it returns an ErrNoService error.
// For each server, it calls the server's Start method and logs the start of the service.
// If any server fails to start, it cancels the context, logs the error, and returns the error.
// It also sets up a signal handler to handle os.Interrupt and syscall.SIGTERM signals.
// If a signal is received, it stops the signal handler, resets the signal handling to the default behavior,
// cancels the context, and returns an error.
// If any server fails while running, it logs the error, cancels the context, and returns the error.
// If the context is cancelled, it logs this event, cancels the context, and returns the context's error.
// After the context is cancelled or a signal is received or a server fails,
// it waits for all servers to stop before returning.
//
// Parameters:
// ctx: The context for the servers. This can be used to cancel the servers or set a deadline.
// logger: The logger to log events.
// servers: The servers to start.
//
// Returns:
// An error if any server fails to start, or if a signal is received, or if any server fails while running.
func StartServers(ctx context.Context, logger *log.ZapEventLogger, servers ...Server) error {
	if len(servers) == 0 {
		return ErrNoService
	}

	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	var dones []Done
	var fails []Fail

	for _, server := range servers {
		logger.Info("starting service: " + server.Name())
		done, fail, err := server.Start(ctx)
		if err != nil {
			cancel()
			logger.Errorw("failed to start service "+server.Name(), "error", err)
			return err
		}
		dones = append(dones, done)
		fails = append(fails, fail)
	}

	anyFail := make(chan error)
	for _, fail := range fails {
		go func(fail Fail) {
			select {
			case v := <-fail:
				select {
				case <-ctx.Done():
				case anyFail <- v:
				}
			case <-ctx.Done():
			default:
			}
		}(fail)
	}

	var err error
	select {
	case <-ctx.Done():
		logger.Debug("context cancelled")
		err = ctx.Err()
		cancel()
	case sig := <-signalChan:
		logger.Warnf("received signal %s", sig.String())
		signal.Stop(signalChan)
		signal.Reset(os.Interrupt, syscall.SIGTERM)
		err = cli.Exit(sig.String(), 130)
		cancel()
	case err = <-anyFail:
		logger.Errorw("service failed", "error", err)
		cancel()
	}

	logger.Info("waiting for services to stop")
	for _, done := range dones {
		<-done
	}

	return err
}
