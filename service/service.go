package service

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pkg/errors"

	"github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
)

type Done = <-chan struct{}

type Fail = <-chan error

// Server is an interface that represents a server that can be started and has a name.
type Server interface {
	// Start Starts the server and returns two channels. The Done channels are closed when the server is done running,
	//	and the Fail channel receives an error if the server fails. The method also returns an error immediately
	//	if the server fails to start. The server runs until the provided context is cancelled.
	Start(ctx context.Context) ([]Done, Fail, error)
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
		dones = append(dones, done...)
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
			}
		}(fail)
	}

	var wg sync.WaitGroup
	var allDone = make(chan struct{})
	for _, done := range dones {
		wg.Add(1)
		go func(done Done) {
			<-done
			wg.Done()
		}(done)
	}
	go func() {
		wg.Wait()
		close(allDone)
	}()

	var err error
	var alreadyDone bool
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
	case <-allDone:
		alreadyDone = true
		logger.Info("all services stopped")
		cancel()
		return err
	}

	if !alreadyDone {
		<-allDone
		logger.Info("all services stopped")
	}
	return err
}
