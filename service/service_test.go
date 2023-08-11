package service

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/ipfs/go-log/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockServer struct {
	mock.Mock
}

func (m *MockServer) Start(ctx context.Context) (Done, Fail, error) {
	args := m.Called(ctx)
	return args.Get(0).(Done), args.Get(1).(Fail), args.Error(2)
}

func (m *MockServer) Name() string {
	return "mock server"
}

func TestStartServers_NoServers(t *testing.T) {
	err := StartServers(context.Background(), log.Logger("test"))
	require.ErrorIs(t, err, ErrNoService)
}

func TestStartServers_Timeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	logger := log.Logger("test")

	// Create a mock server.
	server := new(MockServer)

	done := make(chan struct{})
	// Set up expectations.
	server.On("Start", mock.Anything).Return(Done(done), Fail(nil), nil)

	go func() {
		// Send done once context is cancelled
		<-ctx.Done()
		time.Sleep(time.Second)
		close(done)
	}()
	// Call the function with the mock server.
	err := StartServers(ctx, logger, server)

	// Assert that no error was returned.
	require.ErrorIs(t, err, context.DeadlineExceeded)

	// Assert that the expectations were met.
	server.AssertExpectations(t)
}

func TestStartServers_FailToStart(t *testing.T) {
	ctx := context.Background()
	logger := log.Logger("test")
	server := new(MockServer)

	done := make(chan struct{})
	fail := make(chan error)

	serverError := errors.New("fail to start")

	// Set up expectations.
	server.On("Start", mock.Anything).Return(Done(done), Fail(fail), serverError)

	// Call the function with the mock server.
	err := StartServers(ctx, logger, server)

	// Assert that no error was returned.
	require.ErrorIs(t, err, serverError)

	// Assert that the expectations were met.
	server.AssertExpectations(t)
}

func TestStartServers_ServerError(t *testing.T) {
	ctx := context.Background()
	logger := log.Logger("test")

	// Create a mock server.
	server := new(MockServer)

	// Server is always done
	done := make(chan struct{})
	close(done)

	fail := make(chan error)
	serverError := errors.New("server error")
	go func() {
		fail <- serverError
	}()

	// Set up expectations.
	server.On("Start", mock.Anything).Return(Done(done), Fail(fail), nil)

	// Call the function with the mock server.
	err := StartServers(ctx, logger, server)

	// Assert that no error was returned.
	require.ErrorIs(t, err, serverError)

	// Assert that the expectations were met.
	server.AssertExpectations(t)
}
