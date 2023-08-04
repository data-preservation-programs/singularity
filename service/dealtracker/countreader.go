package dealtracker

import (
	"io"
	"sync"
	"time"
)

// CountingReader is an io.Reader that counts the number of bytes read
type CountingReader struct {
	r io.Reader

	mu        sync.RWMutex // guards n and startTime
	n         int64
	startTime time.Time
}

func NewCountingReader(r io.Reader) *CountingReader {
	return &CountingReader{
		r:         r,
		startTime: time.Now(),
	}
}

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

type Counter interface {
	N() int64
	Speed() float64
}
