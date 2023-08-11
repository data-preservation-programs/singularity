package dealtracker

import (
	"io"
	"sync"
	"time"
)

// CountingReader is an io.Reader that counts the number of bytes read
type CountingReader struct {
	r         io.Reader    // The underlying reader
	mu        sync.RWMutex // Mutex to synchronize access to n and startTime
	n         int64        // The number of bytes read so far
	startTime time.Time    // The time when the counting started
}

func NewCountingReader(r io.Reader) *CountingReader {
	return &CountingReader{
		r:         r,
		startTime: time.Now(),
	}
}

// Read reads data from the underlying reader and updates the byte count.
// It implements the io.Reader interface.
//
// Parameters:
// - p: The byte slice to read data into.
//
// Returns:
// - n: The number of bytes read.
// - err: Any error encountered during the read operation.
func (cr *CountingReader) Read(p []byte) (n int, err error) {
	n, err = cr.r.Read(p)
	cr.mu.Lock()
	defer cr.mu.Unlock()
	cr.n += int64(n)
	return
}

// N returns the number of bytes read so far.
func (cr *CountingReader) N() int64 {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	n := cr.n
	return n
}

// Speed returns the number of bytes read per second
func (cr *CountingReader) Speed() float64 {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	duration := time.Since(cr.startTime).Seconds()
	if duration == 0 {
		return 0
	}
	return float64(cr.n) / duration
}

// Counter represents an interface for counting operations.
//
// The Counter interface defines two methods: N and Speed.
// Implementing types should provide implementations for these methods.
//
// N returns the current byte count.
//
// Speed returns the speed of the counting operation in bytes per second.
type Counter interface {
	N() int64
	Speed() float64
}
