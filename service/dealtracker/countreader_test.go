package dealtracker

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCountingReader_Read(t *testing.T) {
	data := []byte("Hello, World!")
	reader := bytes.NewReader(data)
	countingReader := NewCountingReader(reader)

	require.EqualValues(t, 0, countingReader.N())
	buffer := make([]byte, 5)
	n, err := countingReader.Read(buffer)
	require.NoError(t, err)
	require.EqualValues(t, 5, n)
	require.EqualValues(t, 5, countingReader.N())
}

func TestCountingReader_Speed(t *testing.T) {
	data := []byte("Hello, World!")
	reader := bytes.NewReader(data)
	countingReader := NewCountingReader(reader)

	time.Sleep(1 * time.Second)
	buffer := make([]byte, 5)
	_, err := countingReader.Read(buffer)
	require.NoError(t, err)

	speed := countingReader.Speed()
	require.InDelta(t, 5, speed, 1)
}
