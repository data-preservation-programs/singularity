package service

import (
	"context"
	"testing"
	"time"

	"github.com/cockroachdb/errors"

	"github.com/ipfs/go-log/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockServer struct {
	mock.Mock
	run func(context.Context) error
}

func (m *MockServer) Start(ctx context.Context, exitErr chan<- error) error {
	args := m.Called(ctx, exitErr)
	if m.run == nil {
		exitErr <- nil
	} else {
		exitErr <- m.run(ctx)
	}
	return args.Error(0)
}

func (m *MockServer) Name() string {
	return "mock server"
}

func TestStartServers_NoServers(t *testing.T) {
	err := StartServers(context.Background(), log.Logger("test"))
	require.ErrorIs(t, err, ErrNoService)
}

func TestStartServers_Timeout(t *testing.T) {
	logger := log.Logger("test")

	server := new(MockServer)
	server.run = func(ctx context.Context) error {
		<-ctx.Done()
		return ctx.Err()
	}

	// Set up expectations.
	server.On("Start", mock.Anything, mock.Anything).Return(nil)

	// Start server with canceled context.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := StartServers(ctx, logger, server)

	// Assert that expected error was returned.
	require.ErrorIs(t, err, context.DeadlineExceeded)

	// Assert that the expectations were met.
	server.AssertExpectations(t)
}

func TestStartServers_FailToStart(t *testing.T) {
	logger := log.Logger("test")

	server := new(MockServer)

	serverError := errors.New("fail to start")

	// Set up expectations.
	server.On("Start", mock.Anything, mock.Anything).Return(serverError)

	// Call the function with the mock server.
	err := StartServers(context.Background(), logger, server)

	// Assert that expected error was returned.
	require.ErrorIs(t, err, serverError)

	// Assert that the expectations were met.
	server.AssertExpectations(t)
}

func TestStartServers_ServerError(t *testing.T) {
	logger := log.Logger("test")

	serverError := errors.New("server error")

	server := new(MockServer)
	server.run = func(_ context.Context) error {
		return serverError
	}

	// Set up expectations.
	server.On("Start", mock.Anything, mock.Anything).Return(nil)

	// Call the function with the mock server.
	err := StartServers(context.Background(), logger, server)

	// Assert that no error was returned.
	require.ErrorIs(t, err, serverError)

	// Assert that the expectations were met.
	server.AssertExpectations(t)
}
