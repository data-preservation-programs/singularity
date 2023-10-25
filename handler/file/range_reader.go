package file

import (
	"errors"
	"io"
)

// rangeReader reads data from one individually retrievable file range.
type rangeReader struct {
	// offset is the absolute offset within file where the next read will get
	// data from.
	offset    int64
	reader    io.ReadCloser
	remaining int64
}

func (rr *rangeReader) inRange(pos int64) bool {
	return pos >= rr.offset && pos < rr.offset+rr.remaining
}

func (rr *rangeReader) writeToN(w io.Writer, readLen int64) (int64, error) {
	var read int64
	for readLen > 0 {
		if rr.remaining == 0 {
			return read, io.EOF
		}
		var n int64
		var err error

		if readLen >= rr.remaining {
			// Copy all remaining bytes.
			n, err = io.Copy(w, rr.reader)
		} else {
			// Copy requested number of bytes.
			n, err = io.CopyN(w, rr.reader, readLen)
		}
		if err != nil && !errors.Is(err, io.EOF) {
			return 0, err
		}
		if n == 0 {
			// Must have been EOF.
			rr.remaining = 0
			return read, io.EOF
		}
		rr.offset += n
		rr.remaining -= n
		readLen -= n
		read += n
	}
	return read, nil
}

func (rr *rangeReader) close() error {
	var err error
	if rr.reader != nil {
		if rr.remaining != 0 {
			_, _ = io.Copy(io.Discard, rr.reader)
			rr.remaining = 0
		}
		err = rr.reader.Close()
		rr.reader = nil
	}
	return err
}
