package file

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/handler/handlererror"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var logger = log.Logger("singularity/handler/file")

var (
	ErrNoFileRangeRecord    = errors.New("missing file range record")
	ErrNoJobRecord          = errors.New("missing job record")
	ErrNoFilecoinDeals      = errors.New("no filecoin deals available")
	ErrByteOffsetBeyondFile = errors.New("byte offset would exceed file size")
)

type UnableToServeRangeError struct {
	Start int64
	End   int64
	Err   error
}

func (e UnableToServeRangeError) Unwrap() error {
	return e.Err
}

func (e UnableToServeRangeError) Error() string {
	return fmt.Sprintf("unable to serve byte range %d to %d: %s", e.Start, e.End, e.Err.Error())
}

// RetrieveFileHandler retrieves the actual bytes for a file on disk using a given file ID.
//
// # For now, this function only works if the file remains available in its original source storage
//
// Parameters:
// - ctx: The context for managing timeouts and cancellation.
// - db: The gorm.DB instance for database operations.
// - id: The ID of the file to be retrieved.
//
// Returns:
// - A ReadSeekCloser for the given file
// - the name of the file
// - An error if any issues occur during the database operation, including when the file is not found.
func (DefaultHandler) RetrieveFileHandler(
	ctx context.Context,
	db *gorm.DB,
	filecoinRetriever FilecoinRetriever,
	id uint64,
) (data io.ReadSeekCloser, name string, modTime time.Time, err error) {
	db = db.WithContext(ctx)
	var file model.File
	err = db.Preload("Attachment.Storage").First(&file, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, "", time.Time{}, errors.Wrapf(handlererror.ErrNotFound, "file '%d' does not exist", id)
	}
	if err != nil {
		return nil, "", time.Time{}, errors.WithStack(err)
	}

	rclone, err := storagesystem.NewRCloneHandler(ctx, *file.Attachment.Storage)
	if err != nil {
		return nil, file.FileName(), time.Unix(0, file.LastModifiedNano), errors.WithStack(err)
	}

	seeker, obj, err := storagesystem.Open(rclone, ctx, file.Path)
	if err != nil {
		logger.Infow("no local copy of file, returning filecoin reader", "id", id, "file", file.Path, "fileSize", file.Size, "cid", file.CID)
		// we have no local copy, so we have to return a filecoin based reader
		//nolint:nilerr
		return &filecoinReader{
			ctx:       ctx,
			db:        db,
			retriever: filecoinRetriever,
			size:      file.Size,
			id:        id,
		}, file.FileName(), time.Unix(0, file.LastModifiedNano), nil
	}

	logger.Infow("returning reader for local copy of file", "id", id, "file", file.Path, "fileSize", file.Size, "cid", file.CID)
	return seeker, file.FileName(), obj.ModTime(ctx), nil
}

// io.ReadSeekCloser implementation that reads from remote singularity
type filecoinReader struct {
	ctx       context.Context
	db        *gorm.DB
	retriever FilecoinRetriever
	offset    int64
	size      int64
	id        uint64

	// Reads remaining data from current range.
	rangeReader *rangeReader
}

func (r *filecoinReader) Read(p []byte) (int, error) {
	if r.offset >= r.size {
		return 0, io.EOF
	}

	// Figure out how many bytes to read
	readLen := int64(len(p))
	remainingBytes := r.size - r.offset
	if readLen > remainingBytes {
		readLen = remainingBytes
	}

	buf := bytes.NewBuffer(p)
	buf.Reset()

	n, err := r.writeToN(buf, readLen)
	return int(n), err
}

// WriteTo is implemented in order to directly handle io.Copy operations
// rather than allow small, separate Read operations.
func (r *filecoinReader) WriteTo(w io.Writer) (int64, error) {
	if r.offset >= r.size {
		return 0, io.EOF
	}
	// Read all remaining bytes and write them to w.
	return r.writeToN(w, r.size-r.offset)
}

func (r *filecoinReader) Seek(offset int64, whence int) (int64, error) {
	var newOffset int64

	switch whence {
	case io.SeekStart:
		newOffset = offset
	case io.SeekCurrent:
		newOffset = r.offset + offset
	case io.SeekEnd:
		newOffset = r.size + offset
	default:
		return 0, errors.New("unknown seek mode")
	}

	if newOffset > r.size {
		return 0, ErrByteOffsetBeyondFile
	}

	r.offset = newOffset

	return r.offset, nil
}

func (r *filecoinReader) Close() error {
	var err error
	if r.rangeReader != nil {
		err = r.rangeReader.close()
		r.rangeReader = nil
	}
	return err
}

type deal struct {
	Provider string
}

func findProviders(db *gorm.DB, jobID model.JobID) ([]string, error) {
	var deals []deal
	err := db.Table("deals").Select("distinct provider").
		Joins("JOIN cars ON deals.piece_cid = cars.piece_cid").
		Where("cars.job_id = ? and deals.state IN (?)", jobID, []model.DealState{
			model.DealPublished,
			model.DealActive,
		}).Find(&deals).Error
	if err != nil {
		return nil, err
	}
	providers := make([]string, 0, len(deals))
	for _, deal := range deals {
		providers = append(providers, deal.Provider)
	}
	return providers, nil
}

func findFileRanges(db *gorm.DB, id uint64, startRange int64, endRange int64) ([]model.FileRange, error) {
	var fileRanges []model.FileRange
	err := db.Model(&model.FileRange{}).Where("file_ranges.file_id = ? AND file_ranges.offset < ? AND (file_ranges.offset+file_ranges.length) > ?", id, endRange, startRange).
		Order("file_ranges.offset ASC").Find(&fileRanges).Error
	if err != nil {
		return nil, err
	}
	return fileRanges, nil
}

func (r *filecoinReader) writeToN(w io.Writer, readLen int64) (int64, error) {
	var read int64
	// If there is a rangeReader from the previous read that can be used to
	// continue reading more data, then use it instead of doing another
	// findFileRanges and Retrieve for more reads from this same range.
	if r.rangeReader != nil {
		// If continuing from the previous read, keep reading from this rangeReader.
		if r.offset == r.rangeReader.offset {
			// Reading data leftover from previous read into w.
			n, err := r.rangeReader.writeToN(w, readLen)
			if err != nil && !errors.Is(err, io.EOF) {
				return 0, err
			}
			r.offset += n
			readLen -= n
			read += n
			if readLen == 0 {
				// Read all requested data from leftover in rangeReader.
				return read, nil
			}
			// No more leftover data in rangeReader, but readLen additional bytes
			// still needed. Will read more data from next range(s).
		}
		// No more leftover data in rangeReader, or seek since last read.
		_ = r.rangeReader.close()
		r.rangeReader = nil
	}

	// Get next range(s) to read from.

	// Get all ranges, from current offset, that are covered by readLen.
	fileRanges, err := findFileRanges(r.db, r.id, r.offset, r.offset+readLen)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve file range deals: %w", err)
	}

	var rr *rangeReader

	// Read from each range until readLen bytes read.
	for _, fileRange := range fileRanges {
		if rr != nil {
			_ = rr.close()
			rr = nil
		}
		if readLen == 0 {
			// this shouldn't happen
			logger.Warnw("retrieval reader retrieved file ranges beyond end of range", "fileRangeStart", fileRange.Offset, "fileRangeEnd", fileRange.Offset+fileRange.Length)
			break
		}
		if fileRange.Offset > r.offset {
			return read, UnableToServeRangeError{Start: r.offset, End: fileRange.Offset, Err: ErrNoFileRangeRecord}
		}
		rangeReadLen := readLen
		offsetInRange := r.offset - fileRange.Offset
		remainingRange := fileRange.Length - offsetInRange
		if rangeReadLen > remainingRange {
			rangeReadLen = remainingRange
		}
		// Range starts at fileRange.Offset, has total length fileRange.Length,
		// and has remainingRange bytes left to read. Now read rangeReadLen
		// bytes of the remaining bytes this range.

		if fileRange.JobID == nil {
			return read, UnableToServeRangeError{Start: r.offset, End: r.offset + rangeReadLen, Err: ErrNoJobRecord}
		}
		providers, err := findProviders(r.db, *fileRange.JobID)
		if err != nil || len(providers) == 0 {
			return read, UnableToServeRangeError{Start: r.offset, End: r.offset + rangeReadLen, Err: ErrNoFilecoinDeals}
		}

		// Get a reader that reads until the end of the range.
		rd, err := r.retriever.RetrieveReader(r.ctx, cid.Cid(fileRange.CID), offsetInRange, offsetInRange+remainingRange, providers)
		if err != nil {
			return read, UnableToServeRangeError{
				Start: r.offset,
				End:   r.offset + rangeReadLen,
				Err:   fmt.Errorf("unable to retrieve data from filecoin: %w", err),
			}
		}
		rr = &rangeReader{
			offset:    r.offset,
			reader:    rd,
			remaining: remainingRange,
		}

		// Reading readLen of the remaining bytes in this range.
		n, err := rr.writeToN(w, readLen)
		if err != nil && !errors.Is(err, io.EOF) {
			_ = rr.close()
			return 0, err
		}
		r.offset += n
		readLen -= n
		read += n
	}

	// check for missing file ranges at the end
	if readLen > 0 {
		if rr != nil {
			_ = rr.close()
		}
		return read, UnableToServeRangeError{Start: r.offset, End: r.offset + readLen, Err: ErrNoFileRangeRecord}
	}

	if rr != nil {
		// Some unread data left over in this range. Save it for next read.
		if rr.remaining != 0 {
			// Saving leftover rangeReader with rr.remaining bytes left.
			r.rangeReader = rr
		} else {
			// Leftover rangeReader has 0 bytes remaining.
			_ = rr.close()
		}
	}

	return read, nil
}
