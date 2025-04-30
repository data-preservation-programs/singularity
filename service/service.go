package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cockroachdb/errors"
	"github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)

// Server is an interface that represents a server that can be started and has a name.
type Server interface {
	// Start Starts the server. The method also returns an error immediately if
	// the server fails to start. The server runs until the provided context is
	// cancelled.
	//
	// If an exitErr channel is given, then the server writes an error or nil
	// to the channel when it exits. This channel should be buffered so that
	// every service writing to it perform one non-blocking write.
	Start(ctx context.Context, exitErr chan<- error) error
	// Name returns the name of the server.
	Name() string
}

var ErrNoService = errors.New("no service is running")

// StartServers is responsible for starting the provided servers concurrently and
// managing their lifecycle events such as interrupts or failures. The function
// will ensure that if any server fails or if there is an external interrupt, all
// servers will be gracefully shut down.
//
// Parameters:
//   - ctx: Context that allows for asynchronous task cancellation.
//   - logger: A pointer to a logging instance of type *log.ZapEventLogger.
//   - servers: A variadic parameter that lists all the server instances that need to be started.
//
// Returns:
//   - error: The error encountered during the operation, if any.
//
// The function uses channels to listen for events like server failure or external
// interrupts (like Ctrl+C) and ensures all servers are gracefully terminated
// before exiting.
func StartServers(ctx context.Context, logger *log.ZapEventLogger, servers ...Server) error {
	if len(servers) == 0 {
		return ErrNoService
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	serviceExitErrs := make(chan error, len(servers))
	for _, server := range servers {
		logger.Infow("starting service", "name", server.Name())
		err := server.Start(ctx, serviceExitErrs)
		if err != nil {
			cancel()
			logger.Errorw("failed to start service", "name", server.Name(), "error", err)
			return errors.Wrapf(err, "failed to start service %s", server.Name())
		}
	}

	var err error
	var i int
	running := len(servers)

	for running != 0 {
		select {
		case err = <-serviceExitErrs:
			running--
			i++
			if err != nil {
				logger.Errorw("service failed", "error", err)
			} else {
				logger.Infow("service stopped", "stopped", fmt.Sprintf("%d out of %d", i, len(servers)))
				continue
			}
		case <-ctx.Done():
			logger.Debug("context cancelled")
			err = ctx.Err()
		case sig := <-signalChan:
			logger.Warnf("received signal %s", sig.String())
			signal.Stop(signalChan)
			signal.Reset(os.Interrupt, syscall.SIGTERM)
			err = cli.Exit(sig.String(), 130)
		}
		break
	}

	cancel()
	for running != 0 {
		<-serviceExitErrs
		running--
	}

	logger.Info("all services stopped")
	return errors.WithStack(err)
}
